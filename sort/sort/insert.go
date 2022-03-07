package sort

// InsertSort struct type for sorting
type InsertSort struct{}

// Sort data using the insert sort algorithm
// Use the Sortable interface
func (*InsertSort) Sort(data []int) error {
	for j := 0; j < len(data); j++ {
		idxMin := j
		for i := j + 1; i < len(data); i++ {
			if data[idxMin] > data[i] {
				idxMin = i
			}
		}
		if idxMin != j {
			// go lang usage to swap items
			data[j], data[idxMin] = data[idxMin], data[j]

			// legacy coding
			//			tmp := data[j]
			//			data[j] = data[idxMin]
			//			data[idxMin] = tmp
		}
	}
	return nil
}
