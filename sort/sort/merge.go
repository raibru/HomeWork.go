package sort

// MergeSort struct type for sorting
type MergeSort struct{}

// Sort data using the merge sort algorithm
// Use the Sortable interface
func (ms *MergeSort) Sort(data []int) error {
	ms.mergeSort(data, 0, len(data)-1)
	return nil
}

// sorting the data inside the begin position b and end position e
func (ms *MergeSort) mergeSort(d []int, b int, e int) {
	if b > e {
		return
	}
	ms.mergeSplit(d, b, e)
}

func (ms *MergeSort) mergeSplit(d []int, b int, e int) {
	if b == e {
		return
	}

	m := (b + e) / 2
	ms.mergeSplit(d, b, m)
	ms.mergeSplit(d, m+1, e)
	ms.mergeData(d, b, m, e)
	return
}

func (*MergeSort) mergeData(d []int, b int, m int, e int) {
	ls := 1 + m - b
	if ls < 1 {
		ls = 1
	}
	rs := e - m
	if rs < 1 {
		rs = 1
	}

	var ld = make([]int, ls)
	var rd = make([]int, rs)
	copy(ld[:], d[b:m+1])
	copy(rd[:], d[m+1:e+1])

	i := 0
	j := 0

	for k := b; k <= e; k++ {
		if i < ls && (j >= rs || ld[i] <= rd[j]) {
			d[k] = ld[i]
			i++
		} else {
			d[k] = rd[j]
			j++
		}
	}

	return
}
