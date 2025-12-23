package main

import (
	"os"

	"github.com/jeremyalv/cat-tool/actions"
)

func main() {
	runParams, err := extractRunParams()
	if err != nil {
		panic(err)
	}
	os.Args = append(os.Args[:1], os.Args[2:]...)

	action, err := resolveAction()
	if err != nil {
		panic(err)
	}

	switch action {
	case actions.READ_ARGS:
		processArgs(runParams)
	case actions.READ_STDIN:
		processStdin(runParams)
	case actions.CONCAT:
		processConcat(runParams)
	}
}
