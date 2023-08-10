package heap

type MaxHeap []int

func (h MaxHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MaxHeap) down(parent int) {
	for {
		l := 2*parent + 1
		if l > len(h) {
			break
		}
		j := l
		if r := l + 1; r < len(h) && h.less(l, r) {
			j = r
		}
		if h.less(j, parent) {
			break
		}
		h.swap(parent, j)
		parent = j
	}
}

func (h MaxHeap) up(index int) {
	for {
		parent := (index - 1) / 2
		if parent == index || h.less(index, parent) {
			break
		}
		h.swap(index, parent)
		index = parent
	}
}

func (h MaxHeap) less(i, j int) bool {
	return h[i] < h[j]
}

func (h MaxHeap) Init() {
	n := len(h)

	for i := 2*n - 1; i > 0; i-- {
		h.down(i)
	}
}

func (h *MaxHeap) Pop() int {
	pop := (*h)[0]
	n := len(*h) - 1
	h.swap(0, n)
	*h = (*h)[0:n]
	h.down(0)

	return pop
}

type MaxHeap2 []int

func (heap *MaxHeap2) down(parent int) {
	for {
		l := 2 * parent
		r := 2*parent + 1
		if l > len(*heap) || r > len(*heap) {
			return
		}
		max := l
		if (*heap)[max] < (*heap)[r] {
			max = r
		}
		if (*heap)[max] < (*heap)[parent] {
			break
		}
		heap.swap(max, parent)
		parent = max
	}

}

func (heap *MaxHeap2) heapfiy() {
	for i := len(*heap)/2 - 1; i >= 1; i-- {
		heap.down(i)
	}
}

func NewMaxHeap2(nums []int) *MaxHeap2 {
	maxHeap := MaxHeap2(append([]int{-1}, nums...))
	maxHeap.heapfiy()
	return &maxHeap
}

func (heap *MaxHeap2) swap(i, j int) {
	(*heap)[i], (*heap)[j] = (*heap)[j], (*heap)[i]
}

func (heap *MaxHeap2) Pop() int {
	pop := (*heap)[1]
	heap.swap(1, len(*heap))
	n := len(*heap) - 1
	*heap = (*heap)[0:n]
	heap.down(1)
	return pop
}

// func (h *MaxHeap2) HeapifyMaxHeap2() {
// 	elems := append([]int{-1}, (*h)...)
// }
//
// func (h *MaxHeap2) down(idx int) {
// 	parent := idx / 2
// 	lcli := 2 * parent
// 	rcli := 2*parent + 1
//
// }
