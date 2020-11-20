package merge

import (
	"sort"
)

// Range is a array Containing 2 Ints
type Range [2]int

// List is a collection of Ranges
type List []Range

// sort the list by range
type ByList List
func (l ByList) Len() int           { return len(l) }
func (l ByList) Less(i, j int) bool { return l[i][0] < l[j][0] }
func (l ByList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }


//Merge combines a list of ranges and returns the combined intervals as a new list
// takes a List of Range's and returns aList of Ranges
func Merge(l List) List {
	// list is sorted from smalest starting value to highest starting value
	sort.Sort(ByList(l))
	// define merged List as ml ad append first value to merge list
	var ml List
	ml = append(ml, l[0])

	// define pervious value as p and current value of list as c
	var p Range
	for _, c := range l {
		// pervious is the last value in the ml variable
		p = ml[len(ml)-1]
		// repleace the current end value of the ml variable with the biggest value from p or c
		// if still in the same intervall otherwise append the value to the ml as a new interval
		if c[0] <= p[1] {
			ml[len(ml)-1][1] = max(p[1], c[1])
		} else {
			ml = append(ml, c)
		}
	}

	return ml
}

// max returns the greater number of to ints
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
