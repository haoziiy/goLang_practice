package main

import "fmt"

func isNarcissistic(n int) bool{
	var i, j, k int
	i = n % 10
	j = n / 10 % 10
	k = n / 100 % 10
	sum := i * i * i + j * j * j + k * k * k
	return sum == n
}

func main()  {
	var n int
	fmt.Scanf("%d", &n)
	if isNarcissistic(n) == true {
		fmt.Println(n)
	}
}
