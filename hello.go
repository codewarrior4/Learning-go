package main 

import "fmt"

func main(){

	// working on ranges

	nums := []int{2, 3, 4}

	for _, num := range nums {
		fmt.Println(num)
	}

	for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

	kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
}