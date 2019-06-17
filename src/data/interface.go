package data

import (
	"fmt"
)

type Animal interface {
	eat()
}
type Animal1 interface {
	eat()
}

type Dog struct {
	category string
}
type Cat struct {
	category string
}

// 给结构体定义方法
func (dog Dog) eat() {
	fmt.Println("🐶喜欢吃💩")
}
func (cat Cat) eat() {
	fmt.Println("猫吃罐头")
}

func (cat Cat) miao() {
	fmt.Println("喵")
}
func Interface() {
	var cat = new(Cat)
	cat.eat()
	cat.miao()
	var dog Animal = new(Dog)
	dog.eat()
}
