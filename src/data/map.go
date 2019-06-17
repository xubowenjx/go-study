package data

import (
	"fmt"
)

func TestMap() map[string]string {
	mapper := map[string]string{"wuxi": "taihu"}
	fmt.Println(mapper)
	return mapper
}
