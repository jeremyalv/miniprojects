package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jeremyalv/cat-tool/params"
)

func extractRunParams() ([]params.Param, error) {
	hasParams := len(os.Args) >= 2 && strings.HasPrefix(os.Args[1], string(params.SYMBOL))
	runParams := make([]params.Param, 0)
	if hasParams {
		params, err := resolveParams()
		if err != nil {
			return runParams, fmt.Errorf("failed to resolve params") 
		}
		runParams = params
	}
	return runParams, nil
}