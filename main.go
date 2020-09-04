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
	fmt.Println("(5) Wavelength <-> Wavenumber")
	fmt.Println("(6) Kelvin to thermal Energy")
	fmt.Println("(7) Thermal Energy to Kelvin")
	fmt.Println("Type Quit to close the program")
}

func output(y1 PhysUnit, y2 PhysUnit) {
	fmt.Println(y1.id, y1.value, y1.unit, "was converted to", y2.id, y2.value, y2.unit, ".")
}

func output_multiple(y1 PhysUnitm, y2 PhysUnitm) {
	// output with 2 decimal precision
	fmt.Printf("%s %.2f %s was converted to %s %.2f %s.\n", y1.id, y1.value, y1.unit, y2.id, y2.value, y2.unit)
}

type PhysUnit struct {
	id		string
	value	float64
	unit	string
}

type PhysUnitm struct {
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
			y1 := PhysUnitm{id: "Temperature", value: nums, unit: "F"}
			y2 := PhysUnitm{id: "Temperature", value: ftoCelsius(y1.value), unit: "°C"}
			output_multiple(y1, y2)
		case "2":
			fmt.Println("Type Temperature in Kelvin!")
			nums := myReadln()
			y1 := PhysUnitm{id: "Temperature", value: nums, unit: "K"}
			y2 := PhysUnitm{id: "Temperature", value: kToCelsius(y1.value), unit: "°C"}
			output_multiple(y1, y2)
		case "3":
			fmt.Println("Type Temperature in Celsius!")
			nums := myReadln()
			y1 := PhysUnitm{id: "Temperature", value: nums, unit: "°C"}
			y2 := PhysUnitm{id: "Temperature", value: cToFahrenheit(y1.value), unit: "F"}
			output_multiple(y1, y2)
		case "4":
			fmt.Println("Type Temperature in Celsius!")
			nums := myReadln()
			y1 := PhysUnitm{id: "Temperature", value: nums, unit: "°C"}
			y2 := PhysUnitm{id: "Temperature", value: cToKelvin(y1.value), unit: "K"}
			output_multiple(y1, y2)
		case "5":
			fmt.Println("Type Wavelength/Wavenumber in nm/cm-1!")
			var x float64
			fmt.Scanln(&x)
			y1 := PhysUnit{id: "Wavelength", value: x, unit: "nm"}
			y2 := PhysUnit{id: "Wavenumber", value: waveconvert(x), unit: "cm-1"}
			output(y1, y2)
		case "6":
			fmt.Println("Type Temperature in K!")
			var x float64
			fmt.Scanln(&x)
			fmt.Println(energy_eV(x), "eV")
		case "7":
			fmt.Println("Type Energy in eV!")
			var x float64
			fmt.Scanln(&x)
			fmt.Println(energy_K(x), "K")

		}

	}

}
