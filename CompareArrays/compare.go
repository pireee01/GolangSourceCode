package main

import (
	"fmt"
)

func main() {
	array1 := [...]string{"a", "b", "c"}
	array2 := [...]string{"a", "c", "b"}

	fmt.Println("Array 1 : ", array1)
	fmt.Println("Array 2 : ", array2)

	for i := 0; i < len(array1) && i < len(array2); i++ {
		if array1[i] != array2[i] {
			fmt.Printf("Array berbeda pada indeks %d: %s\n", i, array1[i])
			fmt.Printf("Array berbeda pada indeks %d: %s\n", i, array2[i])
			break
		}
	}
}
