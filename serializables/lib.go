package serializables

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"strings"
	"time"

	"github.com/Lavos/nsid/collections"
	"github.com/eknkc/basex"
)

type Serializable interface {
	Serialize([]byte) string
}

type SerializableFunc func(b []byte) string

func (f SerializableFunc) Serialize(b []byte) string {
	return f(b)
}

func SmallWords(b []byte) string {
	d := make([]string, 8)

	for i, c := range b {
		d[i] = collections.Words[c]
	}

	return strings.Join(d, " ")
}

func StandardBase64(b []byte) string {
	return base64.RawStdEncoding.EncodeToString(b)
}

func BaseX(a string, b []byte) string {
	enc, _ := basex.NewEncoding(a)
	return enc.Encode(b)
}

func Gen() ([]byte, int64) {
	t := time.Now()
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, t.UnixNano())

	return buf.Bytes(), t.UnixNano()
}
