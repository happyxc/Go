package main

func mathSum(a, b int) int {
	return a + b
}

func mathSub(a, b int) int {
	return a - b
}

//MyMath 定义一个函数类型
type MyMath func(int, int) int

//Test 定义的函数类型作为参数使用
func Test(f MyMath, a, b int) int {
	return f(a, b)
}
