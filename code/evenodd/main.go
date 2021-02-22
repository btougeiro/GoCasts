package main

import "fmt"

func main() {
	sliceOfInts := []int{}

	for i := 0; i <= 10; i++ {
		sliceOfInts = append(sliceOfInts, i)
	}

	for _, value := range sliceOfInts {
		if value%2 == 0 {
			fmt.Printf("%d is even\n", value)
		} else {
			fmt.Printf("%d is odd\n", value)
		}
	}
}
