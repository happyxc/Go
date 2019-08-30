package main

import (
	"fmt"
	"reflect"
)

type cat struct {
}

func main() {
	// c := &cat{}
	// typeOfCat := reflect.TypeOf(c)
	// fmt.Println("name: ", typeOfCat.Name()) // 空
	// fmt.Println("kind: ", typeOfCat.Kind())
	x := 3.4
	fmt.Println("type:", reflect.TypeOf(x))
}
