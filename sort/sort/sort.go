package sort

import "strings"

// Sortable interface describe the sot behavior
type Sortable interface {
	Sort([]int) error
}

// Sort sorts a slice of data
func Sort(s Sortable, d []int) error {
	return s.Sort(d)
}

// Int abstact data type
type Int struct {
	value int
}

// LessThan defines how Int are less than another Int type
func (e1 Int) LessThan(e2 Int) bool {
	return e1.value < e2.value
}

// String abstact data type
type String struct {
	value string
}

// LessThan defines how String are less than another String type
func (e1 String) LessThan(e2 String) bool {
	return strings.Compare(e1.value, e2.value) == -1
}
