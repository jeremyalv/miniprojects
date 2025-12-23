package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	BYTES = "c"
	LINES = "l"
	WORDS = "w"
)

func main() {
	var opts string
	var path string
	if len(os.Args) == 2 {
		opts = "-wcl"
		path = os.Args[1]
	} else {
		opts = os.Args[1]
		path = os.Args[2]
	}

	useBytes := strings.Contains(opts, BYTES)
	useLines := strings.Contains(opts, LINES)
	useWords := strings.Contains(opts, WORDS)

	fi, err := os.Open(path)
	if err != nil {
		fmt.Println("error on opening file")
		return
	}
	defer fi.Close()

	fstat, err := fi.Stat()
	if err != nil {
		fmt.Println("error on getting stat")
		return
	}

	name := fstat.Name()
	bytes := fstat.Size()
	
	lines := 0
	words := 0
	sc := bufio.NewScanner(fi)
	for sc.Scan() {
		lines++
		lineWords := strings.Split(sc.Text(), " ")
		words += len(lineWords)
	}
	if err := sc.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	var answer string
	if (useLines) {
		answer += fmt.Sprint(lines, " ")
	}
	if (useWords) {
		answer += fmt.Sprint(words, " ")
	}
	if (useBytes) {
		answer += fmt.Sprint(bytes, " ")
	}

	fmt.Printf("%s %s\n", answer, name)
}
