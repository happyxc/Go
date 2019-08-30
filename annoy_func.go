package main

import "fmt"

func annoyFunc() {
	//	匿名函数，求最大值和最小值
	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(10, 20)
	fmt.Printf("max=%d,min=%d\n", x, y)
}
