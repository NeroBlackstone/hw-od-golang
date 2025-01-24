package mytest

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 库函数积累
// strconv.ParseInt(input, 16, 32) 将十六进制字符串转换为整数
// 第二个参数为0时,默认input带有前缀0x等
func HJ5() {
	// 使用 bufio 读取输入
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	decimal, _ := strconv.ParseInt(input, 16, 32)
	fmt.Println(decimal)

}
