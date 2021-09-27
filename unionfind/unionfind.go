package unionfind

type UnionFind interface {
	Connected(p int, q int) bool
	Union(p int, q int)
	Print()
}
