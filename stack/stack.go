package stack

type Stack interface {
	Push(value int)
	Pop() (int, error)
	IsEmpty() bool
	Size() int
}
