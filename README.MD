# piqad

piqad is a tool to work with Klingon alphabet

## CI Status
[![CircleCI](https://circleci.com/gh/azzz/piqad/tree/master.svg?style=svg)](https://circleci.com/gh/azzz/piqad/tree/master)

# Installation
`go get github.com/azzz/piqad/cmd/piqad`

# Building
You can build application for the most popular platforms manually:
- clone github.com/azzz/piqad
- run `bin/build.sh`

# Usage
`piqad [-from=english] input`

By default, the tool takes input as an English phrase and transliterates to Klingon before outputing UTF-8 codes of Lkingon symbols.

You can pass param `-from-piqad` to interpretate input as a Klingon phrase. In this case, the tool excludes pIqaD tokens from the input and prints related UTF-8 codes and all found characters with their species.
Pay attention that pIqaD is a case-sensetive alphabet.

```
> piqad Uhura
0xF8E5 0xF8D6 0xF8E5 0xF8E1 0xF8D0
Nyota Uhura (Human)
Nyota Uhura ()
Nyota Uhura (Human)
```

```
> piqad -from=piqad Uhura
piqad.Tokenizer: string contains invalid symbols and cannot be parsed well. Stopped at 2-th symbol
```

```
> piqad -from=piqad uHura
0xF8E5 0xF8D6 0xF8E5 0xF8E1 0xF8D0
Nyota Uhura (Human)
Nyota Uhura ()
Nyota Uhura (Human)
```
