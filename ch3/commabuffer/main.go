package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	buf.WriteByte(s[0])

	for i := 1; i < len(s); i++ {
		separator := (len(s) - i) % 3
		if separator == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}

	return buf.String()
}
