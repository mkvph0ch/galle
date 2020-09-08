package conversion

import (
  "fmt"
  "math"
)

func Temperature(nums []float64, inUnit string, outUnit string) []float64{
  // Convert Temperatures
	var result []float64

	if (inUnit == "°C") && (outUnit == "F") {
		// Celsius to Fahrenheit
		for _, v := range nums {
			result = append(result, (v * 9/5) + 32)
		}

	}

	if (inUnit == "°C") && (outUnit == "K") {
		// Celsius to Kelvin
		for _, v := range nums {
			result = append(result, (v + 273.15))
		}

	}

	if (inUnit == "F") && (outUnit == "°C") {
		// Fahrenheit to Celsius
		for _, v := range nums {
			result = append(result, (v-32) * 5/9)
		}

	}

	if (inUnit == "F") && (outUnit == "K") {
		// Fahrenheit to Kelvin
		for _, v := range nums {
			result = append(result, ((v-32) * 5/9) + 273.15)
		}

	}

	if (inUnit == "K") && (outUnit == "°C") {
		// Kelvin to Celsius
		for _, v := range nums {
			result = append(result, v - 273.15)
		}

	}

	if (inUnit == "K") && (outUnit == "F") {
		// Kelvin to Fahrenheit
		for _, v := range nums {
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

func Energy(nums []float64, inUnit string, outUnit string) []float64 {
	// convert Energy <-> Temperature
	var result []float64

	if (inUnit == "K") && (outUnit == "eV") {
		// convert Kelvin K to thermal Energy eV (k_boltzmann * T)
		for _, v := range nums {
			result = append(result, 8.617333262145*math.Pow(10,-5) * v)
			}
	}

	if (inUnit == "eV") && (outUnit == "K") {
		// convert thermal Energy eV to Kelvin (k_boltzmann * T)
		for _, v := range nums {
			result = append(result, v / (8.617333262145*math.Pow(10,-5)))
			}
	}

	return result
}
