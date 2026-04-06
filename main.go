package main

import (
	"fmt"
)

func doCalculation(x int) (s []int) {

	for {
		s = append(s, x)
		if x == 1 {
			break
		}
		if x%2 == 0 {

			x = x / 2
		} else {
			x = x*3 + 1
		}
	}
	return s
}

func main() {
	var base int = 1
	var seed int = 10
	var res2d [][]int

	for ; seed >= base; seed-- { // doCalculation
		res2d = append(res2d, doCalculation(seed))
	}

	for _, row := range res2d {
		fmt.Printf("%v\n", row)
	}
}
