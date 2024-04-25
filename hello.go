package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string
    fmt.Println("uninit:", s, s == nil, len(s) == 0)
	// this print s which is an empty array, next s is empty, and len(s) is 0

	s = make([]string, 3)
	// s=  []string{"a", "b", "c"}
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2]) 

    fmt.Println("len:", len(s))

	s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)


	// you can use the slices.Clone function or copy function to copy a slice
	c := slices.Clone(s)
	// c := make([]string, len(s))
    // copy(c, s)
    fmt.Println("cpy:", c)

	// prints between index 2 and 5 skpping index 5
	l := s[2:5]
    fmt.Println("sl1:", l)

	// prints from 0 to 5 but skips index 5
	l = s[:5]
    fmt.Println("sl2:", l)

	// prints from 2 to the end
	l = s[2:]
    fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    }

	twoD := make([][]int, 4)
    for i := 0; i < 4; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)


}