package main

import (
	"fmt"
	"math"
	"strings"
)

func fahrenheit(a float64) float64 {
	// convert Fahrenheit °F to Celsius °C
	b := (a-32) * 5/9
	return b
}

func kelvin(a float64) float64 {
	// convert Kelvin K to Celsius °C
	b := a - 273.15
	return b
}

func wavelength(a float64) float64 {
	// convert Wavelength nm to Wavenumber cm-1
	b := 1/(a*math.Pow(10, -7))
	return b
}

func wavenumber(a float64) float64  {
	// convert Wavenumber cm-1 to Wavelength nm
	b := 1/(a*math.Pow(10, -7))
	return b
}

func main() {
	fmt.Println("Fahrenheit or Kelvin to Celsius")
	fmt.Println("Wavelength to Wavenumber or Wavenumber to Wavelength")
	var input string
	fmt.Scanln(&input)

	if strings.Compare(input, "Fahrenheit") == 0 {
		fmt.Println("Type Temperature in Fahrenheit!")
		var x float64
		fmt.Scanln(&x)
		fmt.Println(fahrenheit(x))
	}

	if strings.Compare(input, "Kelvin") == 0 {
		fmt.Println("Type Temperature in Kelvin!")
		var x float64
		fmt.Scanln(&x)
		fmt.Println(kelvin(x))
	}

	if strings.Compare(input, "Wavelength") == 0 {
		fmt.Println("Type Wavelength in nm!")
		var x float64
		fmt.Scanln(&x)
		fmt.Println(wavelength(x))
	}

	if strings.Compare(input, "Wavenumber") == 0 {
		fmt.Println("Type Wavenumber in cm-1!")
		var x float64
		fmt.Scanln(&x)
		fmt.Println(wavenumber(x))
	}
}
