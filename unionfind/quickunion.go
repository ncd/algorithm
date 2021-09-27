package unionfind

import (
	"fmt"
	"strconv"
)

type QuickUnion struct {
	id []int
}

func NewQuickUnion(size int) *QuickUnion { // Quick Find
	uf := new(QuickUnion)
	uf.id = make([]int, size)
	for i := 0; i < size; i++ {
		uf.id[i] = i
	}
	return uf
}

func (uf *QuickUnion) root(p int) int {
	for p != uf.id[p] {
		p = uf.id[p]
	}
	return p
}

func (uf *QuickUnion) Connected(p int, q int) bool {
	return uf.root(p) == uf.root(q)
}

func (uf *QuickUnion) Union(p int, q int) {
	proot := uf.root(p)
	qroot := uf.root(q)
	uf.id[proot] = qroot
}

func (uf *QuickUnion) Print() {
	for i := range uf.id {
		fmt.Print(strconv.Itoa(i) + " ")
	}
	fmt.Println()
	for _, v := range uf.id {
		fmt.Print(strconv.Itoa(v) + " ")
	}
	fmt.Println()
}
