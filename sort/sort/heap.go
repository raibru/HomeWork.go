package sort

// HeapSort struct type for sorting
type HeapSort struct{}

// Sort data using the heap sort algorithm
// Use the Sortable interface
func (hs *HeapSort) Sort(data []int) error {
	hs.buildHeap(data)

	for len := len(data); len > 1; len-- {
		hs.removeTop(data, len)
	}
	return nil
}

func (hs *HeapSort) buildHeap(data []int) {
	for i := len(data) / 2; i >= 0; i-- {
		hs.splitHeap(data, i, len(data))
	}
}

func (hs *HeapSort) removeTop(data []int, len int) {
	var lastIdx = len - 1
	data[0], data[lastIdx] = data[lastIdx], data[0]
	hs.splitHeap(data, 0, lastIdx)
}

func (hs *HeapSort) splitHeap(data []int, root int, len int) {
	var max = root
	var l = hs.left(data, root)
	var r = hs.right(data, root)

	if l < len && data[l] > data[max] {
		max = l
	}

	if r < len && data[r] > data[max] {
		max = r
	}

	if max != root {
		data[root], data[max] = data[max], data[root]
		hs.splitHeap(data, max, len)
	}
}

func (*HeapSort) left(data []int, root int) int {
	return (root * 2) + 1
}

func (*HeapSort) right(data []int, root int) int {
	return (root * 2) + 2
}
