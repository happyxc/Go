package main

//Accumulate 闭包实现累加器
func Accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}

}
