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


func waveconvert(a float64) float64  {
	// convert Wavenumber cm-1 <-> Wavelength nm
	if a != 0 {
		b := 1/(a*math.Pow(10, -7))
		return b
	} else {
		fmt.Println("Divison by 0 not possible")
		return 0
	}
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

func options()  {
	fmt.Println("(1) Fahrenheit or (2) Kelvin to Celsius")
	fmt.Println("(3) Wavelength to Wavenumber")
	fmt.Println("(4) Wavenumber to Wavelength")
	fmt.Println("(5) Kelvin to thermal Energy")
	fmt.Println("(6) Thermal Energy to Kelvin")
	fmt.Println("Type Quit to close the program")
}

func main() {
	active := true

	options()

	for active {

		fmt.Println("")
		fmt.Println("Choose a Converter, 'Type Options'")

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
		case "Options":
			options()
		case "1":
			fmt.Println("Type Temperature in Fahrenheit!")
			var x float64
			fmt.Scanln(&x)
			fmt.Println(fahrenheit(x), "ºC")
		case "2":
			fmt.Println("Type Temperature in Kelvin!")
			var x float64
			fmt.Scanln(&x)
			fmt.Println(kelvin(x), "ºC")
		case "3":
			fmt.Println("Type Wavelength in nm!")
			var x float64
			fmt.Scanln(&x)
			fmt.Println(waveconvert(x), "cm-1")
		case "4":
			fmt.Println("Type Wavenumber in cm-1!")
			var x float64
			fmt.Scanln(&x)
			fmt.Println(waveconvert(x), "nm")
		case "5":
			fmt.Println("Type Temperature in K!")
			var x float64
			fmt.Scanln(&x)
			fmt.Println(energy_eV(x), "eV")
		case "6":
			fmt.Println("Type Energy in eV!")
			var x float64
			fmt.Scanln(&x)
			fmt.Println(energy_K(x), "K")

		}

	}

}
