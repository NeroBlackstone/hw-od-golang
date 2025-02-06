package mytest

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func HJ12() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	byteInput := []byte(input)
	for i := 0; i < len(byteInput)/2; i++ {
		byteInput[i], byteInput[len(byteInput)-1-i] = byteInput[len(byteInput)-1-i], byteInput[i]
	}
	fmt.Println(string(byteInput))
}
