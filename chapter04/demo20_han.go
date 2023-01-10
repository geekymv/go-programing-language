package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func main() {
	s := "中(臣)。@#¥%……&"
	IsChineseChar(&s)
	fmt.Println(s)
}

var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")

func filterNonChinese(src *string) {
	strn := ""
	for _, c := range *src {
		if hzRegexp.MatchString(string(c)) {
			strn += string(c)
		}
	}
	*src = strn
}

func IsChineseChar(str *string) {
	strn := ""
	for _, r := range *str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			strn += string(r)
		}
	}
	*str = strn
}
