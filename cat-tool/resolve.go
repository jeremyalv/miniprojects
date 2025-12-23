package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jeremyalv/cat-tool/actions"
	"github.com/jeremyalv/cat-tool/params"
)

func resolveAction() (actions.Action, error) {
	argLength := len(os.Args)

	if argLength == 3 {
		return actions.CONCAT, nil
	}

	if argLength == 1 || argLength == 2 {
		file := os.Stdin
		action, err := resolveReadAction(file)
		
		return action, err
	}

	return actions.NONE, fmt.Errorf("incorrect amount of arguments")
}

func resolveParams() ([]params.Param, error) {
	runParams := make([]params.Param, 0)
	
	if strings.Contains(os.Args[1], string(params.NUMBER_LINES.Identifier)) {
		runParams = append(runParams, params.NUMBER_LINES)
	}

	return runParams, nil
}

func resolveReadAction(file *os.File) (actions.Action, error) {
	fi, err := file.Stat()
	if err != nil {
		return actions.NONE, err
	}

	size := fi.Size()
	if size > 0 {
		return actions.READ_STDIN, nil
	} else {
		return actions.READ_ARGS, nil
	}
}