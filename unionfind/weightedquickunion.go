package unionfind

import (
	"fmt"
	"strconv"
)

type WeightedQuickUnion struct {
	id []int
	sz []int
}

func NewWeightedQuickUnion(size int) *WeightedQuickUnion {
	uf := new(WeightedQuickUnion)
	uf.id = make([]int, size)
	uf.sz = make([]int, size)
	for i := 0; i < size; i++ {
		uf.id[i] = i
		uf.sz[i] = 1
	}
	return uf
}

func (uf *WeightedQuickUnion) root(p int) int {
	for p != uf.id[p] {
		// Pass compression improvement --> Point node to its grandparent
		uf.id[p] = uf.id[uf.id[p]]

		p = uf.id[p]
	}
	return p
}

func (uf *WeightedQuickUnion) Connected(p int, q int) bool {
	return uf.root(p) == uf.root(p)
}

func (uf *WeightedQuickUnion) Union(p int, q int) {
	proot := uf.root(p)
	qroot := uf.root(q)
	if uf.sz[proot] >= uf.sz[qroot] {
		uf.id[qroot] = proot
		uf.sz[proot] += uf.sz[qroot]
	} else {
		uf.id[proot] = qroot
		uf.sz[qroot] += uf.sz[proot]
	}
}

func (uf *WeightedQuickUnion) Print() {
	for i := range uf.id {
		fmt.Print(strconv.Itoa(i) + " ")
	}
	fmt.Println()
	for _, v := range uf.id {
		fmt.Print(strconv.Itoa(v) + " ")
	}
	fmt.Println()
	for _, v := range uf.sz {
		fmt.Print(strconv.Itoa(v) + " ")
	}
	fmt.Println()
}
