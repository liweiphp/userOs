package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CRead() string {
	reader := bufio.NewReader(os.Stdin)
	op, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Print(err)
		return ""
	}
	return strings.TrimSpace(string(op))
}
