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

func cToKelvin(num float64) float64 {
	// convert Celsius °C to Kelvin
	result := num + 273.15
	return result
}

func cToFahrenheit(num float64) float64 {
	// convert Celsius °C to Fahrenheit
	result := (num * 9/5) + 32
	return result
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
	fmt.Println("(3) Celsius to Fahrenheit or (4) to Kelvin")
	fmt.Println("(5) Wavelength to Wavenumber")
	fmt.Println("(6) Wavenumber to Wavelength")
	fmt.Println("(7) Kelvin to thermal Energy")
	fmt.Println("(8) Thermal Energy to Kelvin")
	fmt.Println("Type Quit to close the program")
}

func output(y1 PhysUnit, y2 PhysUnit) {
	fmt.Println(y1.id, y1.value, y1.unit, "was converted to", y2.id, y2.value, y2.unit, ".")
}

type PhysUnit struct {
	id string
	value	float64
	unit string
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
			y1 := PhysUnit{id: "Temperature", value: x, unit: "F"}
			y2 := PhysUnit{id: "Temperature", value: fahrenheit(x), unit: "°C"}
			output(y1, y2)
		case "2":
			fmt.Println("Type Temperature in Kelvin!")
			var x float64
			fmt.Scanln(&x)
			y1 := PhysUnit{id: "Temperature", value: x, unit: "K"}
			y2 := PhysUnit{id: "Temperature", value: kelvin(x), unit: "°C"}
			output(y1, y2)
		case "3":
			fmt.Println("Type Temperature in Celsius!")
			var x float64
			fmt.Scanln(&x)
			y1 := PhysUnit{id: "Temperature", value: x, unit: "°C"}
			y2 := PhysUnit{id: "Temperature", value: cToFahrenheit(x), unit: "F"}
			output(y1, y2)
		case "4":
			fmt.Println("Type Temperature in Celsius!")
			var x float64
			fmt.Scanln(&x)
			y1 := PhysUnit{id: "Temperature", value: x, unit: "°C"}
			y2 := PhysUnit{id: "Temperature", value: cToKelvin(x), unit: "K"}
			output(y1, y2)
		case "5":
			fmt.Println("Type Wavelength in nm!")
			var x float64
			fmt.Scanln(&x)
			y1 := PhysUnit{id: "Wavelength", value: x, unit: "nm"}
			y2 := PhysUnit{id: "Wavenumber", value: waveconvert(x), unit: "cm-1"}
			output(y1, y2)
		case "6":
			fmt.Println("Type Wavenumber in cm-1!")
			var x float64
			fmt.Scanln(&x)
			fmt.Println(waveconvert(x), "nm")
		case "7":
			fmt.Println("Type Temperature in K!")
			var x float64
			fmt.Scanln(&x)
			fmt.Println(energy_eV(x), "eV")
		case "8":
			fmt.Println("Type Energy in eV!")
			var x float64
			fmt.Scanln(&x)
			fmt.Println(energy_K(x), "K")

		}

	}

}
