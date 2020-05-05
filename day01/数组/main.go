package main

import (
	"fmt"
)

func changEle(arr [3]int) {
	arr[0] = 1000
}

func main() {
	var arr1 [5]int
	arr1[1] = 10
	fmt.Println(arr1, arr1[1], arr1[2])

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 34, 4, 5, 6, 7, 78, 8, 98}
	fmt.Println(len(arr3))
	fmt.Println(arr3)
	for i := 0; i < len(arr3)-1; i++ {
		fmt.Println(arr3[i])
	}

	for _, num := range arr3 {
		fmt.Println(num)
	}

	arr4 := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(arr4[1][1])

	arr5 := [3]int{34, 56, 76}
	changEle(arr5)
	fmt.Println(arr5)
}
