package main

import "fmt"

func main() {
	array := make([]int, 0)

	for i := 0; i < 2; i++ {
		array = append(array, 0)
		fmt.Scanf("%d", &array[i])
	}

	if array[0] >= 25 || array[1] <= 40 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
