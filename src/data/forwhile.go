package data

import (
	"fmt"
)

func For() {
	num := []int{2, 3, 4}
	fmt.Println(num)
	for s := 0; s < len(num); s++ {
		fmt.Println(num[s])
	}
	for index, item := range num {
		fmt.Println(item, index)

	}
	var strs [3]string
	strs[1] = "ss"
	strs[2] = "ss1"
	fmt.Println(strs, len(strs))
}
