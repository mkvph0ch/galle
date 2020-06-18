package main

import (
	"fmt"
	//"math"
	"strings"
)

func fahrenheit(a float64) float64 {
	b := (a-32) * 5/9
	return b
}

func kelvin(a float64) float64 {
	b := a - 273.15
	return b
}

func main() {
	fmt.Println("Fahrenheit or Kelvin to Celsius")
	var input string
	fmt.Scanln(&input)
	
	if strings.Compare(input, "Fahrenheit") == 0 {
		fmt.Println("Type Temperature in Fahrenheit!")
		var x float64
		fmt.Scanln(&x)
		fmt.Println(fahrenheit(x))
	}
	
	if strings.Compare(input, "Kelvin") == 0 {
		fmt.Println("Type Temperature in Kelvin!")
		var x float64
		fmt.Scanln(&x)
		fmt.Println(kelvin(x))
	}
}
