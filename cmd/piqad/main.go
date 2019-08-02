package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/azzz/piqad/pkg/piqad"
	"github.com/azzz/piqad/pkg/stapi"
)

type args struct {
	from  string
	input string
}

func main() {
	args, errs := parseArgs()
	if len(errs) > 0 {
		die(errs)
	}

	piqadStr, err := transliterateOrKeep(args)
	if err != nil {
		die(err)
	}

	codes, err := toUTF8Codes(piqadStr)

	if err != nil {
		die(err)
	}

	characters, err := fetchCharacterSpecies(args.input)
	if err != nil {
		die(err)
	}

	fmt.Println(strings.Join(codes, " "))

	cnt := len(characters)
	if cnt == 0 {
		fmt.Println("No characters found with the given name")
	} else {
		fmt.Printf("%d character(s) found with the given name:\n", cnt)
		for _, info := range characters {
			fmt.Printf("%s (%s)\n", info.name, info.specie)
		}
	}
}

type charInfo struct {
	name   string
	specie string
}

func fetchCharacterSpecies(name string) ([]charInfo, error) {
	var (
		conf   *stapi.Configuration = stapi.NewConfiguration()
		client *stapi.APIClient     = stapi.NewAPIClient(conf)
	)
	characters, err := searchCharacters(name, client)

	if err != nil {
		return nil, err
	}

	var uids []string
	for _, char := range characters {
		uids = append(uids, char.Uid)
	}

	var foundChracters []charInfo
	for _, char := range getCharacters(uids, client) {
		info := charInfo{char.Name, speciesToString(char.CharacterSpecies)}
		foundChracters = append(foundChracters, info)
	}

	return foundChracters, nil
}

func die(err interface{}) {
	switch err.(type) {
	case error:
		fmt.Println(err.(error).Error())
	case []error:
		for _, err := range err.([]error) {
			fmt.Println(err.Error())
		}
	}

	os.Exit(2)
}

func transliterateOrKeep(a args) (string, error) {
	switch a.from {
	case "e", "english":
		return piqad.Transliterate(a.input)
	case "p", "piqad":
		return a.input, nil
	}

	return "", errors.New("Source language cannot be reqognized")
}

func toUTF8Codes(str string) (codes []string, err error) {
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

func parseArgs() (args, []error) {
	var errs []error

	from := flag.String(
		"from",
		"english",
		"Sets a source what should be converted to UTF-8 symbols: from \"english\" or from \"piqad\"",
	)

	allowedFromOptions := map[string]bool{
		"english": true,
		"piqad":   true,
		"e":       true,
		"p":       true,
	}

	flag.Parse()
	input := strings.Join(flag.Args(), " ")

	if input == "" {
		errs = append(errs, errors.New("You should provide an input"))
	}

	if _, ok := allowedFromOptions[*from]; !ok {
		errs = append(errs, fmt.Errorf("Invalid -from value. See '%s -h' for details", os.Args[0]))
	}

	return args{from: *from, input: input}, errs
}
