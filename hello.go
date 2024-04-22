package main

import (
	"fmt"
	// "math"
)

func main(){
	// basic type
	// a := 1
	// for a <= 10{
	// 	fmt.Println(a)
	// 	a++
	// }

	// // classic initial
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("Loop", i)
	// }

	// // basic range
	// for i := range 5{
	// 	fmt.Println("Range", i)
	// }

	// // range loop gets to the last item
	// for _, i := range []int{1, 11, 3, 4, 5, 6, 7, 8, 9, 10} {
	// 	fmt.Println("Range", i)
	// }

	// // range loop skips the last item
	// for  i := range []int{1, 11, 3, 4, 5, 6, 7, 8, 9, 10} {
	// 	fmt.Println("Range", i)
	// }

	// with a break
	for {
        fmt.Println("loop")
        break
    }

	// using mod operator and continue
	for n := range 6 {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}