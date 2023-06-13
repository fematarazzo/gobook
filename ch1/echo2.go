// Echo1 exibe seus argumentos de linha de comanda
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println(i)
		fmt.Println(s)
	}
}
