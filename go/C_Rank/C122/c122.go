package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	arr1 := strings.Split(input1, " ")
	all_num, _ := strconv.Atoi(arr1[0])
	condition_num, _ := strconv.Atoi(arr1[1])
	sales_num, _ := strconv.Atoi(arr1[2])

	arr2 := make([]int, all_num)
	for j := 0; j < all_num; j++ {
		input2, _ := reader.ReadString('\n')
		input2 = strings.TrimSpace(input2)
		n, _ := strconv.Atoi(input2)
		arr2[j] = n
	}

	// 問題をちゃんと読んでください！
	lessFunc := func(i, j int) bool {
		return arr2[i] > arr2[j]
	}
	sort.Slice(arr2, lessFunc)

	sum := 0
	sales := false
	if all_num >= condition_num {
		sales = true
	}

	if sales {
		for k := 0; k < all_num-sales_num; k++ {
			sum += arr2[k]
		}
	} else {
		for _, v := range arr2 {
			sum += v
		}
	}

	fmt.Println(sum)
}
