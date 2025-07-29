package main

import (
	"container/list"
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

// ! these three methods are required for working with sort.Sort
func (ps ByName) Len() int {
	return len(ps)
}
func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}
func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	var x list.List
	y := list.New()

	y.PushBack('J')
	y.PushBack('a')
	y.PushBack('m')

	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for element := x.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	kids := []Person{
		{Name: "A", Age: 12},
		{Name: "B", Age: 9},
	}

	sort.Sort(ByName(kids))
	fmt.Println(kids)
}
