package main 

import "fmt"


// basic functions

func addition(a, b int) int{
	return a + b
}

// multiple return values
func vals() (int, int) {
	return 3, 7
}

// variadic functions

func sum(nums ...int) {
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// recursions

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
}

func main(){

	fmt.Println(fact(7))
	fmt.Println(addition(1, 2))
	a, b := vals()
	fmt.Println(a, b)

	_, c := vals()
	fmt.Println(c)

	sum(1, 2)
	sum(1, 2, 3)

	var nums = []int{1, 2, 3, 4}
	sum(nums...)
}