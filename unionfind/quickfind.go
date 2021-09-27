package unionfind

import "fmt"

type QuickFind struct {
	id []int
}

func NewQuickFind(size int) *QuickFind { // Quick Union
	uf := new(QuickFind)
	uf.id = make([]int, size)
	for i := 0; i < size; i++ {
		uf.id[i] = i
	}
	return uf
}

func (uf *QuickFind) Connected(p int, q int) bool {
	return uf.id[p] == uf.id[q]
}

func (uf *QuickFind) Union(p int, q int) {
	pid := uf.id[p]
	qid := uf.id[q]

	for i := range uf.id {
		if uf.id[i] == pid {
			uf.id[i] = qid
		}
	}
}

func (uf *QuickFind) Print() {
	for _, v := range uf.id {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println()
}
