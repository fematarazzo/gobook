// Calls tempconv package
package main

import (
	"fmt"
	"tempconv"
)

func main() {
	fmt.Printf("Brrr! This is Absolute Zero in Celsius %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("Hot! This is Boiling temperature in Farenheit %v\n", tempconv.CToF(tempconv.BoilingC))
	fmt.Printf("Absolutely Hot! This is Boiling temperature in Kelvin %v\n", tempconv.FToK(tempconv.CToF(tempconv.BoilingC)))
}
