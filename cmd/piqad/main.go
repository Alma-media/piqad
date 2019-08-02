package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/azzz/piqad/pkg/piqad"
)

type args struct {
	from  string
	input string
}

func main() {
	args, err := parseArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	codes, err := inputToUTF8Codes(args)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println(strings.Join(codes, " "))
}

func normalizeInput(args args) (piqadString string, err error) {
	switch args.from {
	case "english":
		piqadString, err = piqad.Transliterate(args.input)
		return
	case "piqad":
		piqadString = args.input
		return
	default:
		return piqadString, fmt.Errorf("Invalid -from value. See '%s -h' for details", os.Args[0])
	}
}

func inputToUTF8Codes(args args) (codes []string, err error) {
	str, err := normalizeInput(args)
	if err != nil {
		return
	}

	tokenizer := piqad.NewTokenizer(str)

	for tokenizer.Scan() {
		pchar := tokenizer.Text()
		code, err := piqad.ToUTF8(pchar)

		if err != nil {
			return codes, err
		}

		codes = append(codes, fmt.Sprintf("0x%X", code))
	}

	if err = tokenizer.Err(); err != nil {
		return
	}

	return codes, nil
}

func parseArgs() (args, error) {
	from := flag.String(
		"from",
		"english",
		"Sets a source what should be converted to UTF-8 symbols: from \"english\" or from \"piqad\"",
	)
	flag.Parse()
	input := strings.Join(flag.Args(), " ")

	if input == "" {
		return args{}, errors.New("You should provide an input")
	}
	return args{from: *from, input: input}, nil
}
