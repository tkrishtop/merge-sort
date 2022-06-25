package mergesort

// declare the interface, which is
// lists the methods of some other type
type Sorter interface {
	MergeSort([]int) []int
}

// give a name to a particular function type
// without defining a new one
type SortFunc func([]int) []int

// declare that this function type
// implements this interface
func (s SortFunc) MergeSort(l []int) []int {
	return s(l)
}
