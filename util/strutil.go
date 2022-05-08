package strutil

import (
	"fmt"
	"strings"
)

func init() {
	fmt.Println("util init invoked")
}

func StrToUpper(str string) string {
	return strings.ToUpper(str)
}
