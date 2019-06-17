package data

import (
	"fmt"
	"strconv"
)

const PIE float32 = 3.14 // const

func getValues() (int, string) {
	return 1, "helol"
}
func calcCircle(width float32, height float32) float32 {
	return width * height * PIE
}
func Test() {
	var num int = 1
	strNum := fmt.Sprintf("%d", num)

	var str string = "100"
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d  %s", num+i, strNum+str)
	var stri string = "test"
	stri = "xxx" + stri
	fmt.Println(stri)
	var b bool = true
	fmt.Println(b)

	s1, s2 := getValues()
	fmt.Println(s1, s2)
	fmt.Println(PIE)
	fmt.Println(calcCircle(2, 4))
	ks := [3]int{1, 2, 3}
	fmt.Println("slice", ks[1:])
	for i, s := range ks {
		fmt.Println(i, s)
	}
}
