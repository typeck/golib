package heap


//最大堆
type Heap struct {
	data []int
}

func New(args ...int) *Heap{
	if len(args) == 0 {
		return  &Heap{}
	}
	heap := &Heap{
		data: args,
	}
	heap.heapify()
	return heap
}

//Floyd 建堆算法, 自下而上的下滤o(n)
func (h *Heap)heapify() {
	n := len(h.data)
	//从最后一个有叶子节点的节点开始，下滤
	for i := n/2 - 1; i >= 0; i-- {
		h.sink(i, n)
	}

}

func (h *Heap) Push(x int) {
	h.data = append(h.data, x)
	size := len(h.data)
	if size <= 1 {
		return
	}
	h.swim(size - 1)
}

func (h *Heap)swim(pos int) {
	for {
		p := (pos -1) / 2 //parent
		if p == pos || h.data[p] > h.data[pos] {
			break
		}
		h.data[pos], h.data[p] = h.data[p], h.data[pos]
		pos = p
	}
}

func (h *Heap)Pop() int {
	if len(h.data) == 0 {
		return -1
	}
	top := h.data[0]
	size := len(h.data) 
	h.data[0] = h.data[size-1]
	h.data = h.data[:size-1]
	//弹出一个元素，size--
	h.sink(0, size-1)
	return top
}

func (h *Heap)sink(pos, size int) {
	for {
		lc := 2*pos + 1
		if lc >= size {
			break
		}
		//默认child为左节点
		c := lc
		//如果右节点存在，且右节点大于左节点，则child为右节点
		if rc := lc + 1; rc < size && h.data[rc] > h.data[lc] {
			c = rc
		}
		//即和孩子节点最大的比较
		if h.data[pos] > h.data[c] {
			break
		}
		h.data[pos], h.data[c] = h.data[c], h.data[pos]
		pos = c
	}
}