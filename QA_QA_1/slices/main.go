package main

import "fmt"

func main() {
	var sliceA = []int{0, 1, 2, 3, 4, 123, 345, 22, 43, 12}

	//if 0%2 == 0 {
	//	fmt.Println("Четный")
	//}
	//var strB = []rune("abcde")
	//
	//strNew := append(strB[:1], strB[2:]...)
	//
	//fmt.Println(string(strNew))

	// newSlice := make([]int, 0, len(sliceA))

	for i := 0; i < len(sliceA); i++ {
		if sliceA[i]&1 == 0 { //%2 == 0
			sliceA = append(sliceA[:i], sliceA[i+1:]...)
			i--
		}

	}
	fmt.Println(sliceA)
}

// [min:mid:max]
