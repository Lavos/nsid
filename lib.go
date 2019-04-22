package main

import (
	"time"
	"bytes"
	"encoding/binary"
	"encoding/base64"
	"strings"

	"github.com/eknkc/basex"
)

var (
	base16 = `0123456789abcdef`
	base62 = `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	baseAlpha = `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`
	base85 = `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ.-:+=^!/*?&<>()[]{}@%$#`
	emojis [256]rune
)

func init() {
	emojis[0] = 0x1F004
	emojis[1] = 0x1F0CF
	emojis[2] = 0x1F170
	emojis[3] = 0x1F171
	emojis[4] = 0x1F17E
	emojis[5] = 0x1F17F
	emojis[6] = 0x1F18E
	emojis[7] = 0x1F191
	emojis[8] = 0x1F192
	emojis[9] = 0x1F193
	emojis[10] = 0x1F194
	emojis[11] = 0x1F195
	emojis[12] = 0x1F196
	emojis[13] = 0x1F197
	emojis[14] = 0x1F198
	emojis[15] = 0x1F199
	emojis[16] = 0x1F19A
	emojis[17] = 0x1F1E6
	emojis[18] = 0x1F1E7
	emojis[19] = 0x1F1E8
	emojis[20] = 0x1F1E9
	emojis[21] = 0x1F1EA
	emojis[22] = 0x1F1EB
	emojis[23] = 0x1F1EC
	emojis[24] = 0x1F1ED
	emojis[25] = 0x1F1EE
	emojis[26] = 0x1F1EF
	emojis[27] = 0x1F1F0
	emojis[28] = 0x1F1F1
	emojis[29] = 0x1F1F2
	emojis[30] = 0x1F1F3
	emojis[31] = 0x1F1F4
	emojis[32] = 0x1F1F5
	emojis[33] = 0x1F1F6
	emojis[34] = 0x1F1F7
	emojis[35] = 0x1F1F8
	emojis[36] = 0x1F1F9
	emojis[37] = 0x1F1FA
	emojis[38] = 0x1F1FB
	emojis[39] = 0x1F1FC
	emojis[40] = 0x1F1FD
	emojis[41] = 0x1F1FE
	emojis[42] = 0x1F1FF
	emojis[43] = 0x1F201
	emojis[44] = 0x1F202
	emojis[45] = 0x1F21A
	emojis[46] = 0x1F22F
	emojis[47] = 0x1F232
	emojis[48] = 0x1F233
	emojis[49] = 0x1F234
	emojis[50] = 0x1F235
	emojis[51] = 0x1F236
	emojis[52] = 0x1F237
	emojis[53] = 0x1F238
	emojis[54] = 0x1F239
	emojis[55] = 0x1F23A
	emojis[56] = 0x1F250
	emojis[57] = 0x1F251
	emojis[58] = 0x1F300
	emojis[59] = 0x1F301
	emojis[60] = 0x1F302
	emojis[61] = 0x1F303
	emojis[62] = 0x1F304
	emojis[63] = 0x1F305
	emojis[64] = 0x1F306
	emojis[65] = 0x1F307
	emojis[66] = 0x1F308
	emojis[67] = 0x1F309
	emojis[68] = 0x1F30A
	emojis[69] = 0x1F30B
	emojis[70] = 0x1F30C
	emojis[71] = 0x1F30D
	emojis[72] = 0x1F30E
	emojis[73] = 0x1F30F
	emojis[74] = 0x1F310
	emojis[75] = 0x1F311
	emojis[76] = 0x1F312
	emojis[77] = 0x1F313
	emojis[78] = 0x1F314
	emojis[79] = 0x1F315
	emojis[80] = 0x1F316
	emojis[81] = 0x1F317
	emojis[82] = 0x1F318
	emojis[83] = 0x1F319
	emojis[84] = 0x1F31A
	emojis[85] = 0x1F31B
	emojis[86] = 0x1F31C
	emojis[87] = 0x1F31D
	emojis[88] = 0x1F31E
	emojis[89] = 0x1F31F
	emojis[90] = 0x1F320
	emojis[91] = 0x1F321
	emojis[92] = 0x1F324
	emojis[93] = 0x1F325
	emojis[94] = 0x1F326
	emojis[95] = 0x1F327
	emojis[96] = 0x1F328
	emojis[97] = 0x1F329
	emojis[98] = 0x1F32A
	emojis[99] = 0x1F32B
	emojis[100] = 0x1F32C
	emojis[101] = 0x1F32D
	emojis[102] = 0x1F32E
	emojis[103] = 0x1F32F
	emojis[104] = 0x1F330
	emojis[105] = 0x1F331
	emojis[106] = 0x1F332
	emojis[107] = 0x1F333
	emojis[108] = 0x1F334
	emojis[109] = 0x1F335
	emojis[110] = 0x1F336
	emojis[111] = 0x1F337
	emojis[112] = 0x1F338
	emojis[113] = 0x1F339
	emojis[114] = 0x1F33A
	emojis[115] = 0x1F33B
	emojis[116] = 0x1F33C
	emojis[117] = 0x1F33D
	emojis[118] = 0x1F33E
	emojis[119] = 0x1F33F
	emojis[120] = 0x1F340
	emojis[121] = 0x1F341
	emojis[122] = 0x1F342
	emojis[123] = 0x1F343
	emojis[124] = 0x1F344
	emojis[125] = 0x1F345
	emojis[126] = 0x1F346
	emojis[127] = 0x1F347
	emojis[128] = 0x1F348
	emojis[129] = 0x1F349
	emojis[130] = 0x1F34A
	emojis[131] = 0x1F34B
	emojis[132] = 0x1F34C
	emojis[133] = 0x1F34D
	emojis[134] = 0x1F34E
	emojis[135] = 0x1F34F
	emojis[136] = 0x1F350
	emojis[137] = 0x1F351
	emojis[138] = 0x1F352
	emojis[139] = 0x1F353
	emojis[140] = 0x1F354
	emojis[141] = 0x1F355
	emojis[142] = 0x1F356
	emojis[143] = 0x1F357
	emojis[144] = 0x1F358
	emojis[145] = 0x1F359
	emojis[146] = 0x1F35A
	emojis[147] = 0x1F35B
	emojis[148] = 0x1F35C
	emojis[149] = 0x1F35D
	emojis[150] = 0x1F35E
	emojis[151] = 0x1F35F
	emojis[152] = 0x1F360
	emojis[153] = 0x1F361
	emojis[154] = 0x1F362
	emojis[155] = 0x1F363
	emojis[156] = 0x1F364
	emojis[157] = 0x1F365
	emojis[158] = 0x1F366
	emojis[159] = 0x1F367
	emojis[160] = 0x1F368
	emojis[161] = 0x1F369
	emojis[162] = 0x1F36A
	emojis[163] = 0x1F36B
	emojis[164] = 0x1F36C
	emojis[165] = 0x1F36D
	emojis[166] = 0x1F36E
	emojis[167] = 0x1F36F
	emojis[168] = 0x1F370
	emojis[169] = 0x1F371
	emojis[170] = 0x1F372
	emojis[171] = 0x1F373
	emojis[172] = 0x1F374
	emojis[173] = 0x1F375
	emojis[174] = 0x1F376
	emojis[175] = 0x1F377
	emojis[176] = 0x1F378
	emojis[177] = 0x1F379
	emojis[178] = 0x1F37A
	emojis[179] = 0x1F37B
	emojis[180] = 0x1F37C
	emojis[181] = 0x1F37D
	emojis[182] = 0x1F37E
	emojis[183] = 0x1F37F
	emojis[184] = 0x1F380
	emojis[185] = 0x1F381
	emojis[186] = 0x1F382
	emojis[187] = 0x1F383
	emojis[188] = 0x1F384
	emojis[189] = 0x1F385
	emojis[190] = 0x1F386
	emojis[191] = 0x1F387
	emojis[192] = 0x1F388
	emojis[193] = 0x1F389
	emojis[194] = 0x1F38A
	emojis[195] = 0x1F38B
	emojis[196] = 0x1F38C
	emojis[197] = 0x1F38D
	emojis[198] = 0x1F38E
	emojis[199] = 0x1F38F
	emojis[200] = 0x1F390
	emojis[201] = 0x1F391
	emojis[202] = 0x1F392
	emojis[203] = 0x1F393
	emojis[204] = 0x1F396
	emojis[205] = 0x1F397
	emojis[206] = 0x1F399
	emojis[207] = 0x1F39A
	emojis[208] = 0x1F39B
	emojis[209] = 0x1F39E
	emojis[210] = 0x1F39F
	emojis[211] = 0x1F3A0
	emojis[212] = 0x1F3A1
	emojis[213] = 0x1F3A2
	emojis[214] = 0x1F3A3
	emojis[215] = 0x1F3A4
	emojis[216] = 0x1F3A5
	emojis[217] = 0x1F3A6
	emojis[218] = 0x1F3A7
	emojis[219] = 0x1F3A8
	emojis[220] = 0x1F3A9
	emojis[221] = 0x1F3AA
	emojis[222] = 0x1F3AB
	emojis[223] = 0x1F3AC
	emojis[224] = 0x1F3AD
	emojis[225] = 0x1F3AE
	emojis[226] = 0x1F3AF
	emojis[227] = 0x1F3B0
	emojis[228] = 0x1F3B1
	emojis[229] = 0x1F3B2
	emojis[230] = 0x1F3B3
	emojis[231] = 0x1F3B4
	emojis[232] = 0x1F3B5
	emojis[233] = 0x1F3B6
	emojis[234] = 0x1F3B7
	emojis[235] = 0x1F3B8
	emojis[236] = 0x1F3B9
	emojis[237] = 0x1F3BA
	emojis[238] = 0x1F3BB
	emojis[239] = 0x1F3BC
	emojis[240] = 0x1F3BD
	emojis[241] = 0x1F3BE
	emojis[242] = 0x1F3BF
	emojis[243] = 0x1F3C0
	emojis[244] = 0x1F3C1
	emojis[245] = 0x1F3C2
	emojis[246] = 0x1F3C3
	emojis[247] = 0x1F3C4
	emojis[248] = 0x1F3C5
	emojis[249] = 0x1F3C6
	emojis[250] = 0x1F3C7
	emojis[251] = 0x1F3C8
	emojis[252] = 0x1F3C9
	emojis[253] = 0x1F3CA
	emojis[254] = 0x1F3CB
	emojis[255] = 0x1F3CC
}

type Serializable interface {
	Serialize([]byte) string
}

type SerializableFunc func (b []byte) string

func (f SerializableFunc) Serialize(b []byte) string {
	return f(b)
}

func SmallWords(b []byte) string {
	d := make([]string, 8)

	for i, c := range b {
		d[i] = words[c]
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

func gen() []byte {
	t := time.Now()
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, t.UnixNano())

	return buf.Bytes()
}
