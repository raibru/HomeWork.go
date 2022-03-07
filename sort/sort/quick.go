package sort

// QuickSort struct type for sorting
type QuickSort struct{}

// Sort data using the quick sort algorithm
// Use the Sortable interface
func (qs *QuickSort) Sort(data []int) error {
	qs.quickSort(data, 0, len(data)-1)
	return nil
}

// sorting the data inside the begin position b and end position e
func (qs *QuickSort) quickSort(d []int, b int, e int) {
	if b > e {
		return
	}

	p := qs.partition(d, b, e)
	if p > b {
		qs.quickSort(d, b, p-1)
	}
	if p < e {
		qs.quickSort(d, p+1, e)
	}

	return
}

// partition swap the data around a pivot index and returns the pivot
func (*QuickSort) partition(d []int, b int, e int) int {
	if b > e {
		return e
	}

	piv := d[e]
	j := b
	for i := b; i < e; i++ {
		//for i := range d[b:e] {
		if piv > d[i] {
			d[i], d[j] = d[j], d[i]
			j++
		}
	}
	d[e], d[j] = d[j], d[e]
	return j
}
