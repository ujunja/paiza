package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	s1 = strings.TrimSpace(s1)
	s1_end := s1[len(s1)-1:]

	s2, _ := reader.ReadString('\n')
	s2 = strings.TrimSpace(s2)
	s2_first := s2[:1]
	s2_end := s2[len(s2)-1:]

	if s1_end == s2_first && s2_end != "n" {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}
}
