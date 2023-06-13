// Echo1 exibe seus argumentos de linha de comanda
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(os.Args[0])
	fmt.Println(s)
	seconds := time.Since(start).Seconds()
	fmt.Printf("%.20fs elapsed \n", seconds)
}
