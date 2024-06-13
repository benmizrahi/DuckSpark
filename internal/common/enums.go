package common

type Action int
type Planner int

const (
	TAKE  Action = iota + 1 // EnumIndex = 1
	COUNT                   // EnumIndex = 3
)

func (d Action) EnumIndex() int {
	return int(d)
}

func (d Action) String() string {
	return [...]string{"TAKE", "COUNT"}[d-1]
}

const (
	LIMIT Planner = iota + 1 // EnumIndex = 1
	LOAD                     // EnumIndex = 2
)

func (d Planner) EnumIndex() int {
	return int(d)
}

func (d Planner) String() string {
	return [...]string{"LIMIT", "LOAD"}[d-1]
}
