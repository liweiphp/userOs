package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cfunc func() (bool, error)

func CRead() string {
	reader := bufio.NewReader(os.Stdin)
	op, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Print(err)
		return ""
	}
	return strings.TrimSpace(string(op))
}

func Creturn(cfu Cfunc) bool {
	fmt.Println("=========>>> start =========>>")
	flag, err := cfu()
	if err != nil {
		fmt.Println("系统异常", err)
	}

	fmt.Println("=========>>> end =========>>")
	return flag
}
