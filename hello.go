package main

import (
	"fmt"
	"math"
)

// import "rsc.io/quote"

const s string = "Hello, World!"

func main() {
	fmt.Println(s)
	fmt.Println(math.Pi)
	const n = 500000000

	const d = 30000000000000/n 
	fmt.Println(d)

	fmt.Println(int64(math.Sin(30)))


}