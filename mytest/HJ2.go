package mytest

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 标准库积累
// strings.EqualFold(s, t) 不区分大小写的字符串比较
func HJ2() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')
	// 记得输入字符串一定要去掉首尾空格
	s = strings.TrimSpace(s)

	c, _ := reader.ReadString('\n')
	c = strings.TrimSpace(c)

	sL := strings.ToLower(s)
	cL := strings.ToLower(c)

	count := strings.Count(sL, cL)
	fmt.Println(count)
}
