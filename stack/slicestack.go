package stack

type SliceStack struct {
	value    []int
	capacity int
	size     int
}

func NewSliceStack() *SliceStack {
	s := new(SliceStack)

	return s
}

func (s *SliceStack) Push(value int) {

}

func (s *SliceStack) Pop() int {
	return 0
}

func (s *SliceStack) IsEmpty() bool {
	return s.size == 0
}

func (s *SliceStack) Size() int {
	return s.size
}
