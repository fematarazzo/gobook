package tempconv

// CToF converts Celsius into Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts Fahrenheit into Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts Kelvin into Celsius
func KToC(k Kelvin) Celsius { return Celsius(k + 273.15) }

// FToK converts Fahrenheit into Kelvin
func FToK(f Fahrenheit) Kelvin { return Kelvin((f - 32 + 273.15) * 5 / 9) }
