package main

import (
	"fmt"
	_682 "github.com/arventures/leetcode/682"
)

func main() {
	operations := []string{"5", "2", "C", "D", "+"}
	res := _682.CalPoints(operations)

	// Expected 30
	fmt.Println(res)
}
