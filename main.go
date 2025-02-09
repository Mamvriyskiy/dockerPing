package main

import (
	"fmt"
)

var a int = 1
func main(){
	arrTest := []int{1, 2, 3}

	result := make(map[string]*int, 3)
	for _, v := range arrTest {
		result[fmt.Sprint(v)] = &v
	}

	var str1, str2 string
	for i := 1; i <= 3; i++ {
		str1 = str1 + fmt.Sprint(i)
		str2 = str2 + fmt.Sprint(*result[fmt.Sprint(i)])
	}

	fmt.Println(str1, str2)
}