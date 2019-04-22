package main

import (
	"fmt"
	"os"
	"flag"
)

var (
	suppress_newline = flag.Bool("n", false, "If passed, suppress newline after outputting generated id.")
	format = flag.String("format", "base64", "Encoding format for the generated id.")
)

func main () {
	flag.Parse()
	b := gen()

	switch *format {
	case "phrase":
		fmt.Printf("%s", SmallWords(b))

	case "base16", "hex":
		fmt.Printf("%s", BaseX(base16, b))

	case "base64":
		fmt.Printf("%s", StandardBase64(b))

	case "base62":
		fmt.Printf("%s", BaseX(base62, b))

	case "basealpha":
		fmt.Printf("%s", BaseX(baseAlpha, b))

	case "base85":
		fmt.Printf("%s", BaseX(base85, b))

	default:
		fmt.Fprintf(os.Stderr, "Missing required format.")
		os.Exit(1)
	}

	if *suppress_newline == false {
		fmt.Printf("\n")
	}
}

func test(b []byte) {
	fmt.Printf("Phrase:\t\t%s\n", SmallWords(b))
	fmt.Printf("Base16:\t\t%s\n", BaseX(base16, b))
	fmt.Printf("Base64:\t\t%s\n", StandardBase64(b))
	fmt.Printf("Base62:\t\t%s\n", BaseX(base62, b))
	fmt.Printf("BaseAlpha:\t%s\n", BaseX(baseAlpha, b))
	fmt.Printf("Base85:\t\t%s\n", BaseX(base85, b))
}
