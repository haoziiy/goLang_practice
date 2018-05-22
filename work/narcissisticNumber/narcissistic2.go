package main

import (
	"fmt"
	"strconv"
)

func isNarcissistic(str string) bool {
	var result int
	for i := 0; i < len(str) ; i++  {
		num := int(str[i] - '0')
		result += (num * num * num)
	}

	number, error := strconv.Atoi(str)
	if error != nil {
		fmt.Println("can not convert %s to int\n", str)
		return false
	}

	if result == number {
		fmt.Printf("%d\n", result)
		return true
	} else {
		return false
	}
}

func main()  {
	var str string
	fmt.Scanf("%s", &str)
	if isNarcissistic(str) == true {
		fmt.Println(str + " is a narcissistic number")
	}
}