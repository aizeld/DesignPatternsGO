// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Sorter interface {
	Sort([]int) []int
}

type BubbleSort struct{}

func (bs *BubbleSort) Sort(arr []int) []int { return arr }

type InsertionSort struct{}

func (is *InsertionSort) Sort(arr []int) []int { return arr }

type Context struct {
	sorter Sorter
}

func (c *Context) SetSorter(sorter Sorter) {
	c.sorter = sorter
}

func (c *Context) ExecuteSort(arr []int) []int { return c.sorter.Sort(arr) }

func main() {
	arr := []int{4, 4, 3, 6, 6, 2, 23, 5, 6, 7, 67, 8, 89, 5, 3, 2, 2, 34, 65, 65, 2}
	bubbleSort := &BubbleSort{}
	insertionSort := &InsertionSort{}

	context := &Context{}

	context.SetSorter(bubbleSort)
	fmt.Println("Bubble : ", context.ExecuteSort(arr))

	context.SetSorter(insertionSort)
	fmt.Println("Insertion : ", context.ExecuteSort(arr))

}
