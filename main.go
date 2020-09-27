package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"encoding/csv"

	cv "galle/conversion"
	op "galle/options"
	rl "galle/readln"
)

type PhysUnit struct {
	// struct PhysUnit for all physical Units
	id    string
	value []float64
	unit  string
}

func toCSV(y2 PhysUnit) {
	recordFile, err := os.Create("./out.csv")

	if err != nil {
		fmt.Println(err)
	}

	newStr := []string{}

	for _, v := range y2.value {
		newStr = append(newStr, strconv.FormatFloat(v, 'f', -1, 64))
	}

	fmt.Println("This is newStr", newStr)

	w := csv.NewWriter(recordFile)
	w.Write(newStr)
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

func output(y1 PhysUnit, y2 PhysUnit) {
	fmt.Println(y1.id, y1.value, y1.unit, "was converted to", y2.id, y2.value, y2.unit, ".")

	toCSV(y2)
}

func outputPrecision(y1 PhysUnit, y2 PhysUnit, precision int) {
	// output with 2 decimal precision
	fmt.Printf("%s %.*f %s was converted to %s %.*f %s.\n", y1.id, precision, y1.value, y1.unit, y2.id, precision, y2.value, y2.unit)
}

func main() {
	active := true

	m := op.Options()

	for active {

		fmt.Printf("\nChoose a Converter, Type \"Options\" \n")

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
		case m["F->C"]:
			fmt.Println("Type Temperature in Fahrenheit!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "F"}
			y2 := PhysUnit{id: "Temperature", value: cv.Temperature(y1.value, y1.unit, "°C"), unit: "°C"}
			output(y1, y2)
		case m["F->K"]:
			fmt.Println("Type Temperature in Fahrenheit!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "F"}
			y2 := PhysUnit{id: "Temperature", value: cv.Temperature(y1.value, y1.unit, "K"), unit: "K"}
			output(y1, y2)
		case m["K->C"]:
			fmt.Println("Type Temperature in Kelvin!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "K"}
			y2 := PhysUnit{id: "Temperature", value: cv.Temperature(y1.value, y1.unit, "°C"), unit: "°C"}
			output(y1, y2)
		case m["K->F"]:
			fmt.Println("Type Temperature in Kelvin!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "K"}
			y2 := PhysUnit{id: "Temperature", value: cv.Temperature(y1.value, y1.unit, "F"), unit: "F"}
			output(y1, y2)
		case m["C->F"]:
			fmt.Println("Type Temperature in Celsius!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "°C"}
			y2 := PhysUnit{id: "Temperature", value: cv.Temperature(y1.value, y1.unit, "F"), unit: "F"}
			output(y1, y2)
		case m["C->K"]:
			fmt.Println("Type Temperature in Celsius!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "°C"}
			y2 := PhysUnit{id: "Temperature", value: cv.Temperature(y1.value, y1.unit, "K"), unit: "K"}
			output(y1, y2)
		case m["nm->cm-1"]:
			fmt.Println("Type Wavelength in nm!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Wavelength", value: nums, unit: "nm"}
			y2 := PhysUnit{id: "Wavenumber", value: cv.Waveconvert(y1.value), unit: "cm-1"}
			output(y1, y2)
		case m["cm-1->nm"]:
			fmt.Println("Type Wavenumber in cm-1!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Wavenumber", value: nums, unit: "cm-1"}
			y2 := PhysUnit{id: "Wavelength", value: cv.Waveconvert(y1.value), unit: "nm"}
			output(y1, y2)
		case m["K->eV"]:
			fmt.Println("Type Temperature in K!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Temperature", value: nums, unit: "K"}
			y2 := PhysUnit{id: "Energy", value: cv.Energy(y1.value, y1.unit, "eV"), unit: "eV"}
			output(y1, y2)
		case m["eV->K"]:
			fmt.Println("Type Energy in eV!")
			nums := rl.Readln()
			y1 := PhysUnit{id: "Energy", value: nums, unit: "eV"}
			y2 := PhysUnit{id: "Temperature", value: cv.Energy(y1.value, y1.unit, "K"), unit: "K"}
			output(y1, y2)
		}

	}

}
