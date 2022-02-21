package tempconv

// CToF converts Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c - KelvinToCelsiusDifference)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k + KelvinToCelsiusDifference)
}

func FToK(f Fahrenheit) Kelvin {
	c := FToC(f)
	return CToK(c)
}

func KToF(k Kelvin) Fahrenheit {
	c := KToC(k)
	return CToF(c)
}
