package protos

type Action int

const (
	TAKE  Action = iota + 1 // EnumIndex = 1
	LIMIT                   // EnumIndex = 2
	COUNT                   // EnumIndex = 3
	LOAD
)

func (d Action) EnumIndex() int {
	return int(d)
}

func (d Action) String() string {
	return [...]string{"Take", "LIMIT", "COUNT", "LOAD"}[d-1]
}
