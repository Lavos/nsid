package main

import (
	"fmt"
	"os"
	"flag"

	"github.com/Lavos/nsid/collections"
	"github.com/Lavos/nsid/serializables"
)

var (
	suppress_newline = flag.Bool("n", false, "If passed, suppress newline after outputting generated id.")
	format = flag.String("format", "base64", "Encoding format for the generated id. Formats: all unixnano phrase base16 hex base62 base64 basealpha base85")
)

func main () {
	flag.Parse()
	b, n := serializables.Gen()

	switch *format {
	case "unixnano":
		fmt.Printf("%d", n)

	case "phrase":
		fmt.Printf("%s", serializables.SmallWords(b))

	case "base16", "hex":
		fmt.Printf("%s", serializables.BaseX(collections.Base16, b))

	case "base64":
		fmt.Printf("%s", serializables.StandardBase64(b))

	case "base62":
		fmt.Printf("%s", serializables.BaseX(collections.Base62, b))

	case "basealpha":
		fmt.Printf("%s", serializables.BaseX(collections.BaseAlpha, b))

	case "base85":
		fmt.Printf("%s", serializables.BaseX(collections.Base85, b))

	case "all":
		fmt.Printf("UnixNano:\t%d\n", n)
		fmt.Printf("Phrase:\t\t%s\n", serializables.SmallWords(b))
		fmt.Printf("Base16:\t\t%s\n", serializables.BaseX(collections.Base16, b))
		fmt.Printf("Base64:\t\t%s\n", serializables.StandardBase64(b))
		fmt.Printf("Base62:\t\t%s\n", serializables.BaseX(collections.Base62, b))
		fmt.Printf("BaseAlpha:\t%s\n", serializables.BaseX(collections.BaseAlpha, b))
		fmt.Printf("Base85:\t\t%s\n", serializables.BaseX(collections.Base85, b))

	default:
		fmt.Fprintf(os.Stderr, "Missing required format.")
		os.Exit(1)
	}

	if *suppress_newline == false {
		fmt.Printf("\n")
	}
}
