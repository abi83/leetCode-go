package main

import (
	"fmt"
	"leetCode/Alghorithm"
)

func main() {
	var a = []int{1, 2}
	Alghorithm.RotateArrayCpuEffective(a, 5)
	fmt.Println(a)
}
