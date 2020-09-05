package main

import (
	"fmt"
	"math"
	"os"
	"time"
	"bufio"
	"strconv"
	"strings"
)

func ftoCelsius(nums []float64) []float64 {
	// convert Fahrenheit °F to Celsius °C
	var result []float64

	for _, v := range nums {
		result = append(result, (v-32) * 5/9)
		}
	return result
}

func kToCelsius(nums []float64) []float64 {
	// convert Kelvin K to Celsius °C
	var result []float64

	for _, v := range nums {
		result = append(result, v - 273.15)
		}
	return result
}

func cToKelvin(nums []float64) []float64 {
	// convert Celsius °C to Kelvin
	var result []float64

	for _, v := range nums {
		result = append(result, v + 273.15)
		}
	return result
}

func cToFahrenheit(nums []float64) []float64 {
	// convert Celsius °C to Fahrenheit
	var result []float64

	for _, v := range nums {
		result = append(result, (v * 9/5) + 32)
		}
	return result
}

func waveconvert(nums []float64) []float64  {
	// convert Wavenumber cm-1 <-> Wavelength nm
	var result []float64

	for _, v := range nums {

		if v != 0 {
			result = append(result, 1/(v*math.Pow(10, -7)))
		} else {
			fmt.Println("Divison by 0 not possible")
		}

		}
	return result
}

func energy_eV(nums []float64) []float64  {
	// convert Kelvin K to thermal Energy eV (k_boltzmann * T)
	var result []float64

	for _, v := range nums {
		result = append(result, 8.617333262145*math.Pow(10,-5) * v)
		}
	return result
}

func energy_K(nums []float64) []float64  {
	// convert thermal Energy eV to Kelvin (k_boltzmann * T)
	var result []float64

	for _, v := range nums {
		result = append(result, v / 8.617333262145*math.Pow(10,-5))
		}
	return result
}

func options()  {
	options := []string{"Fahrenheit to Celsius",
											"Kelvin to Celsius",
											"Celsius to Fahrenheit",
											"Celsius to Kelvin",
											"Wavelength to Wavenumber",
											"Wavenumber to Wavelength",
											"Kelvin to thermal Energy",
											"Thermal Energy to Kelvin",
											"Type Quit to exit"}

	for i, v := range options {
		if i+1 < len(options) {
			fmt.Println(i+1, v)
		} else {
			fmt.Println(v)
		}

	}
}

func output(y1 PhysUnit, y2 PhysUnit) {
	fmt.Println(y1.id, y1.value, y1.unit, "was converted to", y2.id, y2.value, y2.unit, ".")
}

func output_precision(y1 PhysUnit, y2 PhysUnit) {
	// output with 2 decimal precision
	fmt.Printf("%s %.2f %s was converted to %s %.2f %s.\n", y1.id, y1.value, y1.unit, y2.id, y2.value, y2.unit)
}

type PhysUnit struct {
	id		string
	value	[]float64
	unit	string
}

func myReadln() []float64 {
	// read multiple input arguments
	reader := bufio.NewReader(os.Stdin)
	var nums []float64

	for {

		fmt.Print("Enter number: ")

		text, err := reader.ReadString('\n')
		if err != nil {
			// error handling
		}
		text = strings.TrimSpace(text)

		num, err := strconv.ParseFloat(text, 64)
		if err != nil {
			// error handling
			break
		}

		nums = append(nums, num)
		}

		return nums
}

func main() {
	active := true

	options()

	for active {

		fmt.Printf("\nChoose a Converter, Type 'Options'\n")

		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)

		switch input {
		default:
			fmt.Println("Please choose a conversion from the list!")
		case "quit":
			fmt.Println("Program will be closed")
			active = false
			time.Sleep(time.Millisecond * 2000)
			os.Exit(1)
		case "options":
			options()
		case "1":
			fmt.Println("Type Temperature in Fahrenheit!")
			nums := myReadln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "F"}
			y2 := PhysUnit{id: "Temperature", value: ftoCelsius(y1.value), unit: "°C"}
			output(y1, y2)
		case "2":
			fmt.Println("Type Temperature in Kelvin!")
			nums := myReadln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "K"}
			y2 := PhysUnit{id: "Temperature", value: kToCelsius(y1.value), unit: "°C"}
			output(y1, y2)
		case "3":
			fmt.Println("Type Temperature in Celsius!")
			nums := myReadln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "°C"}
			y2 := PhysUnit{id: "Temperature", value: cToFahrenheit(y1.value), unit: "F"}
			output(y1, y2)
		case "4":
			fmt.Println("Type Temperature in Celsius!")
			nums := myReadln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "°C"}
			y2 := PhysUnit{id: "Temperature", value: cToKelvin(y1.value), unit: "K"}
			output(y1, y2)
		case "5":
			fmt.Println("Type Wavelength in nm!")
			nums := myReadln()
			y1 := PhysUnit{id: "Wavelength", value: nums, unit: "nm"}
			y2 := PhysUnit{id: "Wavenumber", value: waveconvert(y1.value), unit: "cm-1"}
			output(y1, y2)
		case "6":
			fmt.Println("Type Wavenumber in cm-1!")
			nums := myReadln()
			y1 := PhysUnit{id: "Wavenumber", value: nums, unit: "cm-1"}
			y2 := PhysUnit{id: "Wavelength", value: waveconvert(y1.value), unit: "nm"}
			output(y1, y2)
		case "7":
			fmt.Println("Type Temperature in K!")
			nums := myReadln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "K"}
			y2 := PhysUnit{id: "Energy", value: energy_eV(y1.value), unit: "eV"}
			output(y1, y2)
		case "8":
			fmt.Println("Type Energy in eV!")
			nums := myReadln()
			y1 := PhysUnit{id: "Energy", value: nums, unit: "eV"}
			y2 := PhysUnit{id: "Temperature", value: energy_K(y1.value), unit: "K"}
			output(y1, y2)
		}

	}

}
