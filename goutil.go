package goutil

import (
	"fmt"
)

func SayHello() {
	fmt.Println("hello qzp")
}

func FindType(i interface{}) {
	switch x := i.(type) {
	case int:
		fmt.Println(x, "is int")
	case string:
		fmt.Println(x, "is string")
	case nil:
		fmt.Println(x, "is nil")
	default:
		fmt.Println(x, "not type matched")
	}
}
