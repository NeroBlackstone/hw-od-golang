package mytest

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 库函数积累
// strings.Repeat(s, n) 重复字符串
func HJ4() {
	// 使用 bufio 读取输入
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')

	// 去掉首尾多余的空格和换行符
	s = strings.TrimSpace(s)

	// 处理字符串，按照每 8 个字符分组
	n := len(s)
	for i := 0; i < n; i += 8 {
		if i+8 <= n {
			// 如果可以完整取出 8 个字符，直接输出
			fmt.Println(s[i : i+8])
		} else {
			// 如果不满 8 个字符，用 '0' 补齐
			fmt.Println(s[i:] + strings.Repeat("0", 8-(n-i)))
		}
	}
}
