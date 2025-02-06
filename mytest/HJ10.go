package mytest

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 考set去重
func HJ10() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	seen := make(map[byte]bool)
	for i := 0; i < len(input); i++ {
		seen[input[i]] = true
	}
	fmt.Println(len(seen))
}
