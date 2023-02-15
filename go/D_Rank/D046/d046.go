package main

import "fmt"

func main() {
	var num int
	var tmp int
	for i := 0; i < 3; i++ {
		fmt.Scanf("%d", &tmp)
		if num == 0 || num < tmp {
			num = tmp
		}
	}
	fmt.Println(num)
}
