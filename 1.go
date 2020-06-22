package main

import "fmt"

// 234168
func main()  {
	var sum int = 0

	for i := 1; i <= 1000; i++ {
		if ( i % 3 ==0 || i % 5 == 0 ) {
			sum += i
		}
	}

	fmt.Println( sum )
}