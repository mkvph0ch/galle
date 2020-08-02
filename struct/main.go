package main

import (
	"fmt"
  "math"
)
// Objektbeschreibung Viereck. Großbuchstabe = public, Kleinbuchstabe = private
type Viereck struct {
	a, b	float64
}
// Objektbeschreibung Dreieck
type Dreieck struct {
  a, b, c	float64
}
// Objektbeschreibung Kreis
type Kreis struct {
  r float64
}
// Methode von Dreieck
func (t Dreieck) area() float64 {
	return t.a * t.b * t.c
}
// Methode von Viereck
func (t Viereck) area() float64 {
	return t.a * t.b
}
// Methode von Viereck
func (t Kreis) area() float64 {
  return math.Pi * t.r * t.r
}
// interface Geometry
type Geometry interface {
  area() float64
}

func measure(g Geometry) {
    fmt.Println("This is the Object", g)
    fmt.Println("This is the area of the Object", g.area())
}

// Objekt Physikalische Größe
type PhysUnit struct {
	id	string
	value	float64
	unit	string
	category	string
}
// PhysConst erbt Beschreibung von PhysUnit
type PhysConst struct {
	PhysUnit
	popularity	float64
}

func main() {
	// Definition von Vier- und Dreieck
	x1 := Viereck{a: 4.0, b: 8.0}
	x2 := Dreieck{a: 5.0, b: 9.0, c: 10.0}
  x3 := Kreis{r: 4.00}

	fmt.Println("This is the area of the Object", x1, x1.area())
	fmt.Println("This is the area of the Object", x2, x2.area())

  // ODER
  measure(x1)
  measure(x2)
  measure(x3)

	// Definition der Boltzmann-Konstante
	y1 := PhysConst{
		PhysUnit: PhysUnit{
			id: "kB",
			value: 1.38e-23,
			unit: "J/K",
			category: "Energy",
		 },
		 popularity: 30,
	}
	// Zugriff auf key-value Paare von Boltzmann-Konstante
	fmt.Println(y1.popularity) // von PhysConst
	fmt.Println(y1.id) // von PhysUnit
	fmt.Println(y1.value) // von PhysUnit

  // Definition einer Spannung
  v1 := PhysUnit{"Voltage", 3, "V", "Electromagnetism"}
  // Definition einer Spannung
	v2 := PhysUnit{"Voltage", 4, "V", "Electromagnetism"}
  // Definition einer Spannung, Addition von s und a
	result := PhysUnit{v1.id, v1.value + v2.value, v1.unit, v1.category}

	fmt.Printf("The object %s has a value of %.2f and the unit %s\n", v1.id, v1.value, v1.unit)
	fmt.Printf("The object %s has a value of %.2f and the unit %s\n", v2.id, v2.value, v2.unit)
	fmt.Printf("The object %s has a value of %.2f and the unit %s\n", result.id, result.value, result.unit)
}
