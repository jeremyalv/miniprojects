package params

import (
	"fmt"
	"strings"
)

type NumberLines struct {
	Identifier string
}

func (p NumberLines) Process(data string) string {
	trimmedData := strings.TrimSuffix(data, "\n")
	lines := strings.Split(trimmedData, "\n")
	processed := make([]string, 0)
	
	for i, line := range lines {
		processed = append(processed, fmt.Sprintf("%d %s", i+1, line))
	}

	result := strings.Join(processed, "\n")
	return result
}
