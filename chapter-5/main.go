package main

import "fmt"

func main() {
	arr()
	fmt.Println("---Array function triggered---")
	slice()
	fmt.Println("---Slice function triggered---")
	sliceAppend()
	fmt.Println("---SliceAppend function triggered---")
	sliceCopy()
	fmt.Println("---SliceCopy function triggered---")
	Map()
	fmt.Println("---Map function triggered---")
	complexMap()
	fmt.Println("---ComplexMap function triggered---")
	findMinNm()
	fmt.Println("---FindMinNum function triggered---")
}

// Arrays
func arr() {
	x := [5]float64{34, 35, 36, 37, 38}
	var total float64 = 0
	for _, value := range x { // The _ is index that we didn't need here
		total += value
	}
	fmt.Println(total / float64(len(x)))
}

// Slices
func slice() {
	arr := [5]float64{1, 2, 3, 4, 5}
	//x := make([]float64, 5)
	x := arr[1:4]
	y := arr[:] // shallow copy of whole array => y = arr[0 : len(x)]
	fmt.Println(x, y)
}

// Slice , Append
func sliceAppend() {
	//x := make([]int, 0)  // This type of declaring
	var x []int // Or this type of declaring
	x = append(x, 4, 5)
	y := append(x, 2, 1, 45, 8)
	fmt.Println(x, y)
}

// Slice , Copy
func sliceCopy() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := make([]int, 2)
	copy(slice2, slice1) // The content of second arg(slice) is copied in the first arg(slice)
	fmt.Println(slice1, slice2)
	fmt.Println(slice1[len(slice1)-1])
}

// Map in golang
func Map() {
	//elements := make(map[string]string) // The first type is for "key" & the second is for value
	//elements["H"] = "Hydrogen"
	//elements["He"] = "Helium"
	//elements["Li"] = "Lithium"
	//elements["F"] = "Fluorine"
	//elements["Ne"] = "Neon"
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(elements["Li"])
	fmt.Println(elements)
	fmt.Println(len(elements))
	delete(elements, "Ne")
	fmt.Println(len(elements))
	fmt.Println(elements)
	if name, ok := elements["F"]; ok {
		fmt.Println(name, ok)
	}
}

// A little more complex maps
func complexMap() {
	elements := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
		"F": {
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": {
			"name":  "Neon",
			"state": "gas",
		},
	}
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}

// Function for finding the minimum number in array
func findMinNm() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, -12,
		37, 34, 9, 27,
		17, 18, 97, 83,
	}
	temp := x[0]
	for _, val := range x {
		if temp > val {
			temp = val
		}
	}
	fmt.Printf("The minimum number in array is %d\n", temp)
}
