package mytest

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 标准库函数积累
// bufio.NewReader(os.Stdin) 读取一行字符串
// strings.TrimSpace(input) 去掉首尾空格
// strings.Split(input, " ") 按字符串分割
func HJ1() {
	// 读取一行字符串
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	// 去掉首尾空格（边界条件记得处理）
	input = strings.TrimSpace(input)
	words := strings.Split(input, " ")
	lastWord := words[len(words)-1]
	fmt.Println(len(lastWord))
}
