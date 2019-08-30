package main

import "fmt"

//Animal interface
type Animal interface {
	Move()
}

//Bird struct
type Bird struct {
	Name string
}

//Move :Bird struct实现Animal接口
func (b *Bird) Move() {
	fmt.Println("bird fly...")
}

//Fish struct
type Fish struct {
	Name string
}

//Move :Fish struct实现Animal接口
func (f *Fish) Move() {
	fmt.Println("fish swim...")
}

//Factory 函数传递一个name，返回Animal接口
func Factory(name string) Animal {
	switch name {
	case "bird":
		return &Bird{}
	case "fish":
		return &Fish{}
	default:
		return nil
	}
}
