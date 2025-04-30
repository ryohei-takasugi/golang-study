package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()

	// cast
	var j interface{} = "hello2"
	s := j.(string)
	fmt.Println(s)

	s, ok := j.(string)
	fmt.Println(s, ok)

	f, ok := j.(float64)
	fmt.Println(f, ok)

	doSwithType(0)
	doSwithType("hello")
	doSwithType(true)
}

func doSwithType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("unknown")
	}
}
