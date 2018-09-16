package sorting

import (
	"fmt"
	"sort"
)

type Extremity struct {
	min int
	max int64
}

type ExtremityList []Extremity

func (c ExtremityList) Len() int           { return len(c) }
func (c ExtremityList) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ExtremityList) Less(i, j int) bool { return c[i].min < c[j].min }

func Solution(A []int) int {
	list := make([]Extremity, len(A))
	for i, val := range A {
		list[i].min = i - val
		list[i].max = int64(i) + int64(val)
	}

	sort.Sort(ExtremityList(list))

	count := 0

	fmt.Printf("%v\n", list)
	for c := 0; c < len(list)-1; c++ {
		ff := sort.Search(len(list), func(i int) bool { return int64(list[i].min) > list[c].max })
		count += ff - (c+1)
	}

	if count > 10000000 {
		return -1
	}

	return count
}

func main() {
	A := []int{1,1,1}
	fmt.Println(Solution(A))

	B := []int{-4, -1, 0, 0, 2, 5}
	ff := sort.Search(len(B), func(i int) bool { return B[i] >= 1 })
	fmt.Println(ff)
}