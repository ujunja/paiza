package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 100点ではない！
	var reader = bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	a := strings.Split(s, "")
	var num = a[0]
	for i := 0; i < len(a)-1; i++ {
		if a[i] != num && a[i] != " " {
			fmt.Println("No")
			break
		} else {
			if i == len(a)-2 {
				fmt.Println("Yes")
			}
		}
	}

}
