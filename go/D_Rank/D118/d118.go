package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	arr := make([]string, 2)
	for i := 0; i < 2; i++ {
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		arr[i] = s
	}
	n, _ := strconv.Atoi(arr[1])
	if arr[0] == "S" {
		fmt.Println(1925 + n)
	} else {
		fmt.Println(1988 + n)
	}
}
