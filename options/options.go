package options

import (
	"fmt"
	"strconv"
)

func Options() map[string]string {
	// show converting options
	m := make(map[string]string)

	options := [][]string{{"Fahrenheit to Celsius", "F->C"},
		{"Fahrenheit to Kelvin", "F->K"},
		{"Kelvin to Celsius", "K->C"},
		{"Kelvin to Fahrenheit", "K->F"},
		{"Celsius to Fahrenheit", "C->F"},
		{"Celsius to Kelvin", "C->K"},
		{"Wavelength to Wavenumber", "nm->cm-1"},
		{"Wavenumber to Wavelength", "cm-1->nm"},
		{"Kelvin to thermal Energy", "K->eV"},
		{"Thermal Energy to Kelvin", "eV->K"},
		{"Type Quit to exit", "quit"}}

	for i, v := range options {
		if i+1 < len(options) {
			fmt.Printf("%d. %s\n", i+1, v[0])
			m[v[1]] = strconv.Itoa(i + 1)
		} else {
			fmt.Println(v[0])
			m[v[1]] = strconv.Itoa(0)
		}
	}

	return m
}
