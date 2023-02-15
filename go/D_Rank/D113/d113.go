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
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	arr := strings.Split(s, ":")
	n, _ := strconv.Atoi(arr[0])
	n += 16
	if n >= 24 {
		n -= 24
	}
	fmt.Print(n, ":", arr[1])
}
