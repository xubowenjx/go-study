package data

import (
	"fmt"
)

func Cursor() {
	var c *int
	var k = 20
	c = &k
	fmt.Println(c, *c)
}
