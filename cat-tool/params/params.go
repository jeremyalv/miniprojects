package params

type Param interface {
	Process(data string) string
}

const (
	SYMBOL string = "-"
)

var (
	NUMBER_LINES = NumberLines{
		Identifier: "n",
	}
)
