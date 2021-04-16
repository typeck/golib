package uf_set

//并查集 union find set

// 并查集（Union-find Sets）是一种非常精巧而实用的数据结构，它主要用于处理一些不相交集合的合并问题。
// 一些常见的用途有求连通子图、求最小生成树的 Kruskal 算法和求最近公共祖先（Least Common Ancestors, LCA）等。

// https://leetcode-cn.com/problems/redundant-connection/description/
// https://zhuanlan.zhihu.com/p/93647900/
type UninoFindSet struct {
	parent []int
	rank []int
}

func NewUnionFindSet(n int) *UninoFindSet{
	var parent = make([]int, n)
	var rank = make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i 
		rank[i] = 1
	}
	return &UninoFindSet{
		parent: parent,
		rank: rank,
	}
}

func (uf *UninoFindSet) Find(x int) int{
	for uf.parent[x] != x {
		//并查集中只关心根节点
		//路径压缩 : x.parent = x.parent.parent
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}
	return x
}

//递归写法
func (uf *UninoFindSet) FindRec(x int) int {
	if uf.parent[x] == x {
		return x
	}
	uf.parent[x] = uf.FindRec(uf.parent[x]) //父节点设为根节点
	return uf.parent[x] //返回父节点
}

func (uf *UninoFindSet) Merge(i, j int) {	
	x := uf.Find(i)
	y := uf.Find(j)

	// 按秩合并：每次合并都把深度较小的集合合并在深度较大的集合下面
	if uf.rank[x] < uf.rank[y] {
		uf.parent[x] = y
	}else {
		uf.parent[y] = x
	}
	//如果深度相同且根节点不同，则新的根节点的深度+1
	if uf.rank[x] == uf.rank[y] && x != y {
		uf.rank[y]++
	}
}