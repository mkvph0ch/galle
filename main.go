package main

import (
	"fmt"
	"os"
	"time"
	"strings"

	cv "galle/conversion"
	op "galle/options"
	rl "galle/readln"
)

type PhysUnit struct {
	// struct PhysUnit for all physical Units
	id		string
	value	[]float64
	unit	string
}

func output(y1 PhysUnit, y2 PhysUnit) {
	fmt.Println(y1.id, y1.value, y1.unit, "was converted to", y2.id, y2.value, y2.unit, ".")
}

func outputPrecision(y1 PhysUnit, y2 PhysUnit, precision int) {
	// output with 2 decimal precision
	fmt.Printf("%s %.*f %s was converted to %s %.*f %s.\n", y1.id, precision, y1.value, y1.unit, y2.id, precision, y2.value, y2.unit)
}

func main() {
	active := true

	m := op.Options()

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
			m = op.Options()
		case m["Fahrenheit to Celsius"]:
			fmt.Println("Type Temperature in Fahrenheit!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "F"}
			y2 := PhysUnit{id: "Temperature", value: cv.Temperature(y1.value, y1.unit, "°C"), unit: "°C"}
			output(y1, y2)
		case m["Fahrenheit to Kelvin"]:
			fmt.Println("Type Temperature in Fahrenheit!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "F"}
			y2 := PhysUnit{id: "Temperature", value: cv.Temperature(y1.value, y1.unit, "K"), unit: "K"}
			output(y1, y2)
		case m["Kelvin to Celsius"]:
			fmt.Println("Type Temperature in Kelvin!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "K"}
			y2 := PhysUnit{id: "Temperature", value: cv.Temperature(y1.value, y1.unit, "°C"), unit: "°C"}
			output(y1, y2)
		case m["Kelvin to Fahrenheit"]:
			fmt.Println("Type Temperature in Kelvin!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "K"}
			y2 := PhysUnit{id: "Temperature", value: cv.Temperature(y1.value, y1.unit, "F"), unit: "F"}
			output(y1, y2)
		case m["Celsius to Fahrenheit"]:
			fmt.Println("Type Temperature in Celsius!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "°C"}
			y2 := PhysUnit{id: "Temperature", value: cv.Temperature(y1.value, y1.unit, "F"), unit: "F"}
			output(y1, y2)
		case m["Celsius to Kelvin"]:
			fmt.Println("Type Temperature in Celsius!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "°C"}
			y2 := PhysUnit{id: "Temperature", value: cv.Temperature(y1.value, y1.unit, "K"), unit: "K"}
			output(y1, y2)
		case m["Wavelength to Wavenumber"]:
			fmt.Println("Type Wavelength in nm!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Wavelength", value: nums, unit: "nm"}
			y2 := PhysUnit{id: "Wavenumber", value: cv.Waveconvert(y1.value), unit: "cm-1"}
			output(y1, y2)
		case m["Wavenumber to Wavelength"]:
			fmt.Println("Type Wavenumber in cm-1!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Wavenumber", value: nums, unit: "cm-1"}
			y2 := PhysUnit{id: "Wavelength", value: cv.Waveconvert(y1.value), unit: "nm"}
			output(y1, y2)
		case m["Kelvin to thermal Energy"]:
			fmt.Println("Type Temperature in K!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "K"}
			y2 := PhysUnit{id: "Energy", value: cv.Energy(y1.value, y1.unit, "eV"), unit: "eV"}
			output(y1, y2)
		case m["Thermal Energy to Kelvin"]:
			fmt.Println("Type Energy in eV!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Energy", value: nums, unit: "eV"}
			y2 := PhysUnit{id: "Temperature", value: cv.Energy(y1.value, y1.unit, "K"), unit: "K"}
			output(y1, y2)
		}

	}

}
