package main

import "fmt"

func main() {
	xs := []float64{23, 24, 25, 26}
	xs2 := []float64{1, 2, 3, 4, 5}
	slc := []int{1, 2, 3, 4}
	fmt.Println(average(xs))
	fmt.Println(average2(xs2))
	x, y := f()
	fmt.Printf("x: %d and y: %d \n", x, y)
	fmt.Printf("the sum of numbers: %d \n", add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37))
	fmt.Println(add(slc...)) // slice ellipsis

	//closure
	z := 0
	increment := func() int {
		z++
		return z
	}
	fmt.Println("closure: ", increment())
	fmt.Println("closure: ", increment())
	nextEven := makeEvenGenerator()
	fmt.Println("closure-2: ", nextEven())
	fmt.Println("closure-2: ", nextEven())
	fmt.Println("closure-2: ", nextEven())

	fmt.Println("Factorial)", factorial(5))

	defer second()
	first()

	// Pointer
	q := 5
	zero(q)
	fmt.Println(q) // 5
	zero2(&q)
	fmt.Println(q) // 0

	// new Pointer
	mathPiNum := new(float64)
	PI(mathPiNum)
	fmt.Println(*mathPiNum) // 3.1415
}

// Functions
func average(arr []float64) float64 {
	total := 0.0
	for _, v := range arr {
		total += v
	}
	return total / float64(len(arr))
}
func average2(arr []float64) (answer float64) {
	total := 0.0
	for _, v := range arr {
		total += v
	}
	answer = total / float64(len(arr))
	return
}
func f() (int, int) {
	return 5, 6
} // Multiple return values
func add(args ...int) int {
	total := 0
	for _, val := range args {
		total += val
	}
	return total
}
func makeEvenGenerator() func() int {
	i := 0
	return func() int {
		ret := i
		i += 2
		return ret
	}
} // Closure
func factorial(index int) int {
	if index == 0 {
		return 1
	}
	return index * factorial(index-1)
}
func first() {
	fmt.Println("First")
}
func second() {
	fmt.Println("Second")
}

// Pointer
func zero(q int) {
	q = 0
} // Without Pointer
func zero2(qPtr *int) {
	*qPtr = 0
} // With Pointer
func PI(piNumber *float64) {
	*piNumber = 3.1415
}