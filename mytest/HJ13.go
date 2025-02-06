package mytest

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func HJ13() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	splitInput := strings.Split(input, " ")
	result := ""
	for i := 0; i < len(splitInput); i++ {
		result += splitInput[len(splitInput)-1-i] + " "
	}
	result = strings.TrimSpace(result)
	fmt.Println(result)
}
