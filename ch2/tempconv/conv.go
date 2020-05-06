package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius(5 * (f - 32) / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }
