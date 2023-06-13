// Echo1 exibe seus argumentos de linha de comanda
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	seconds := time.Since(start).Seconds()
	fmt.Printf("%.20fs elapsed \n", seconds)
}
