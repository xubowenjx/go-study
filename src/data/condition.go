package data

import (
	"fmt"
)

func Condition() {
	var num int = 5
	if num < 3 {
		fmt.Println("if true")
	} else {
		fmt.Println("more than 3")
	}
}
