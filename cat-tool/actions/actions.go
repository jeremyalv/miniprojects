package actions

type Action int

const (
	NONE Action = iota
	READ_ARGS
	READ_STDIN
	CONCAT
)
