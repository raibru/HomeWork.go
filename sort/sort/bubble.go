package sort

// BubbleSort struct type for sorting
type BubbleSort struct{}

// Sort data using the bubble sort algorithm
// Use the Sortable interface
func (bs *BubbleSort) Sort(data []int) error {
	for j := len(data); j > 1; j-- {
		if !bs.bubbleSort(data, j) {
			break
		}
	}
	return nil
}

func (*BubbleSort) bubbleSort(data []int, last int) bool {
	isSwap := false
	for i := 1; i < last; i++ {
		if data[i] < data[i-1] {
			// go lang usage to swap items
			data[i], data[i-1] = data[i-1], data[i]
			isSwap = true
		}
	}
	return isSwap
}
