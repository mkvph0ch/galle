package main

import (
	"fmt"
	"math"
	//"strings"
	"os"
	"time"
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

func energy_eV(a float64) float64  {
	// convert Kelvin K to thermal Energy eV (k_boltzmann * T)
	b := 8.617333262145*math.Pow(10,-5) * a
	return b
}

func energy_K(a float64) float64  {
	// convert thermal Energy eV to Kelvin (k_boltzmann * T)
	b := a / 8.617333262145*math.Pow(10,-5)
	return b
}

func main() {
	active := true

	for active {
		fmt.Println("Fahrenheit or Kelvin to Celsius")
		fmt.Println("Wavelength to Wavenumber")
		fmt.Println("Wavenumber to Wavelength")
		fmt.Println("Kelvin to thermal Energy")
		fmt.Println("Thermal Energy to Kelvin")
		fmt.Println("Type Quit to close the program")
		var input string
		fmt.Scanln(&input)

		switch input {
		default:
			fmt.Println("Please choose a conversion from the list!")
		case "Quit":
			fmt.Println("Program will be closed")
			active = false
			time.Sleep(time.Millisecond * 2000)
			os.Exit(1)
		case "Fahrenheit":
			fmt.Println("Type Temperature in Fahrenheit!")
			var x float64
			fmt.Scanln(&x)
			fmt.Println(fahrenheit(x))
		case "Kelvin":
			fmt.Println("Type Temperature in Kelvin!")
			var x float64
			fmt.Scanln(&x)
			fmt.Println(kelvin(x))
		case "Wavelength":
			fmt.Println("Type Wavelength in nm!")
			var x float64
			fmt.Scanln(&x)
			fmt.Println(wavelength(x))
		case "Wavenumber":
			fmt.Println("Type Wavenumber in cm-1!")
			var x float64
			fmt.Scanln(&x)
			fmt.Println(wavenumber(x))
		case "Energy":
			fmt.Println("Type Temperature in K!")
			var x float64
			fmt.Scanln(&x)
			fmt.Println(energy_eV(x))
		case "Temperature":
			fmt.Println("Type Energy in eV!")
			var x float64
			fmt.Scanln(&x)
			fmt.Println(energy_K(x))

		}
		
	}

}
