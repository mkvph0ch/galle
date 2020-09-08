package options

import (
  "strconv"
  "fmt"
)

func Options() map[string]string {

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
