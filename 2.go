package main

import "fmt"

// 4613732
func main() {
	fibNum_arr := []int{ 1, 2 }
	var fibNum int = 0
	var evenNum int = 0

	for i := 3; fibNum < 4000000; i++ {
		fibNum = 0

		fibNum += fibNum_arr[ len( fibNum_arr ) - 1 ] + fibNum_arr[ len( fibNum_arr ) - 2 ]

		fibNum_arr = append( fibNum_arr, fibNum )
	}

	for _, v := range fibNum_arr {
		if v % 2 == 0 {
			evenNum += v
		}
	}

	fmt.Println( evenNum )
}

