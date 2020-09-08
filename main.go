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

func Temperature(y PhysUnit, s string) []float64{

	var result []float64

	if (y.unit == "°C") && (s == "F") {
		// Celsius to Fahrenheit
		for _, v := range y.value {
			result = append(result, (v * 9/5) + 32)
		}

	}

	if (y.unit == "°C") && (s == "K") {
		// Celsius to Kelvin
		for _, v := range y.value {
			result = append(result, (v + 273.15))
		}

	}

	if (y.unit == "F") && (s == "°C") {
		// Fahrenheit to Celsius
		for _, v := range y.value {
			result = append(result, (v-32) * 5/9)
		}

	}

	if (y.unit == "F") && (s == "K") {
		// Fahrenheit to Kelvin
		for _, v := range y.value {
			result = append(result, ((v-32) * 5/9) + 273.15)
		}

	}

	if (y.unit == "K") && (s == "°C") {
		// Kelvin to Celsius
		for _, v := range y.value {
			result = append(result, v - 273.15)
		}

	}

	if (y.unit == "K") && (s == "F") {
		// Kelvin to Fahrenheit
		for _, v := range y.value {
			result = append(result, (v - 273.15) * 9/5 + 32)
		}

	}

	return result
}

func Waveconvert(nums []float64) []float64  {
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

func Energy(y PhysUnit, s string) []float64 {
	//
	var result []float64

	if (y.unit == "K") && (s == "eV") {
		// convert Kelvin K to thermal Energy eV (k_boltzmann * T)
		for _, v := range y.value {
			result = append(result, 8.617333262145*math.Pow(10,-5) * v)
			}
	}

	if (y.unit == "eV") && (s == "K") {
		// convert thermal Energy eV to Kelvin (k_boltzmann * T)
		for _, v := range y.value {
			result = append(result, v / (8.617333262145*math.Pow(10,-5)))
			}
	}

	return result
}

func options() map[string]string {

	m := make(map[string]string)

	options := []string{"Fahrenheit to Celsius",
											"Fahrenheit to Kelvin",
											"Kelvin to Celsius",
											"Kelvin to Fahrenheit",
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
			m[v] = strconv.Itoa(i+1)
		} else {
			fmt.Println(v)
			m[v] = strconv.Itoa(0)
		}
	}

	return m
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
			break
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

	m := options()

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
			m = options()
		case m["Fahrenheit to Celsius"]:
			fmt.Println("Type Temperature in Fahrenheit!")
			nums := myReadln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "F"}
			y2 := PhysUnit{id: "Temperature", value: Temperature(y1, "°C"), unit: "°C"}
			output(y1, y2)
		case m["Fahrenheit to Kelvin"]:
			fmt.Println("Type Temperature in Fahrenheit!")
			nums := myReadln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "F"}
			y2 := PhysUnit{id: "Temperature", value: Temperature(y1, "K"), unit: "K"}
			output(y1, y2)
		case m["Kelvin to Celsius"]:
			fmt.Println("Type Temperature in Kelvin!")
			nums := myReadln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "K"}
			y2 := PhysUnit{id: "Temperature", value: Temperature(y1, "°C"), unit: "°C"}
			output(y1, y2)
		case m["Kelvin to Fahrenheit"]:
			fmt.Println("Type Temperature in Kelvin!")
			nums := myReadln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "K"}
			y2 := PhysUnit{id: "Temperature", value: Temperature(y1, "F"), unit: "F"}
			output(y1, y2)
		case m["Celsius to Fahrenheit"]:
			fmt.Println("Type Temperature in Celsius!")
			nums := myReadln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "°C"}
			y2 := PhysUnit{id: "Temperature", value: Temperature(y1, "F"), unit: "F"}
			output(y1, y2)
		case m["Celsius to Kelvin"]:
			fmt.Println("Type Temperature in Celsius!")
			nums := myReadln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "°C"}
			y2 := PhysUnit{id: "Temperature", value: Temperature(y1, "K"), unit: "K"}
			output(y1, y2)
		case m["Wavelength to Wavenumber"]:
			fmt.Println("Type Wavelength in nm!")
			nums := myReadln()
			y1 := PhysUnit{id: "Wavelength", value: nums, unit: "nm"}
			y2 := PhysUnit{id: "Wavenumber", value: Waveconvert(y1.value), unit: "cm-1"}
			output(y1, y2)
		case m["Wavenumber to Wavelength"]:
			fmt.Println("Type Wavenumber in cm-1!")
			nums := myReadln()
			y1 := PhysUnit{id: "Wavenumber", value: nums, unit: "cm-1"}
			y2 := PhysUnit{id: "Wavelength", value: Waveconvert(y1.value), unit: "nm"}
			output(y1, y2)
		case m["Kelvin to thermal Energy"]:
			fmt.Println("Type Temperature in K!")
			nums := myReadln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "K"}
			y2 := PhysUnit{id: "Energy", value: Energy(y1, "eV"), unit: "eV"}
			output(y1, y2)
		case m["Thermal Energy to Kelvin"]:
			fmt.Println("Type Energy in eV!")
			nums := myReadln()
			y1 := PhysUnit{id: "Energy", value: nums, unit: "eV"}
			y2 := PhysUnit{id: "Temperature", value: Energy(y1, "K"), unit: "K"}
			output(y1, y2)
		}

	}

}
