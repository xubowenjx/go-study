package user

import (
	"fmt"
	"godemo/src/data"
)

type Person struct {
	name    string
	address string
	age     int
}

//把这里改成 main 就可以 go run main.go 执行
func Test() {
	data.Test()
	data.Condition()
	data.For()
	data.Cursor()
	p := Person{name: "xubw", address: "wuxi", age: 23}
	fmt.Println(p)
	var oc *Person
	oc = &p
	fmt.Println(oc, *oc)
	mp := data.TestMap()
	for key, value := range mp {
		fmt.Println(key, value)
	}

	var i = 2
	value := string(i)
	fmt.Println(value)
	data.Interface()
}
