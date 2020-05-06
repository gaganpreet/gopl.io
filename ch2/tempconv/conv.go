package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius(5 * (f - 32) / 9) }
