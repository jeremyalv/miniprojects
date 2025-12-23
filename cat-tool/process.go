package main

import (
	"fmt"
	"io"
	"os"

	"github.com/jeremyalv/cat-tool/params"
)

func processArgs(params []params.Param) {
	fmt.Printf("args: %s", os.Args)
	path := os.Args[1]
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)	
	}
	resultText := string(dat)

	result := processParams(resultText, params)
	fmt.Printf("%s", result)
}

func processStdin(params []params.Param) {
	file := os.Stdin
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	resultText := string(bytes)

	result := processParams(resultText, params)
	fmt.Printf("%s", result)
}

func processConcat(params []params.Param) {
	file1 := os.Args[1]
	file2 := os.Args[2]

	dat1, err := os.ReadFile(file1)
	if err != nil {
		panic(err)
	}
	dat2, err := os.ReadFile(file2)
	if err != nil {
		panic(err)
	}
	resultText := string(dat1) + string(dat2)

	result := processParams(resultText, params)
	fmt.Printf("%s", result)
}

func processParams(text string, params []params.Param) string {
	for _, param := range params {
		text = param.Process(text)
	}

	return text
}
