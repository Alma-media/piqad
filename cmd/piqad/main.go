package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/azzz/piqad/pkg/piqad"
)

type args struct {
	source string
	input  string
}

func main() {
	args := parseArgs()
	var piqadString string

	switch args.source {
	case "english":
		var err error
		piqadString, err = piqad.Transliterate(args.input)

		if err != nil {
			die(err.Error())
		}
	case "piqad":
		piqadString = args.input
	default:
		die(fmt.Sprintf("Invalid -from value. See '%s -h' for details", os.Args[0]))
	}

	codes := toUtf8Codes(piqadString)
	fmt.Println(strings.Join(codes, " "))
}

func parseArgs() args {
	source := flag.String(
		"from",
		"english",
		"Sets a source what should be converted to UTF-8 symbols: from \"english\" or from \"piqad\"",
	)
	flag.Parse()
	input := strings.Join(flag.Args(), " ")

	if input == "" {
		die("You should provide an input")
	}
	return args{source: *source, input: input}
}

func die(err string) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}

func toUtf8Codes(piqadString string) []string {
	codes := make([]string, 0)
	tokenizer := piqad.NewTokenizer(piqadString)

	for tokenizer.Scan() {
		pchar := tokenizer.Text()

		if code, err := piqad.ToUTF8(pchar); err != nil {
			die(err.Error())
		} else {
			codes = append(codes, fmt.Sprintf("0x%X", code))
		}
	}

	if err := tokenizer.Err(); err != nil {
		die(err.Error())
	}

	return codes
}
