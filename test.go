package main

import (
	"fmt"
	//"math"
	//"strings"
)
//kommentar
func main() {
  var mystring string
  mystring = "Birne"

  var mystring2 string
  mystring2 = "Apfel"

  if mystring == mystring2 {
    fmt.Println("Identical")
  }

  if mystring != mystring2 {
    fmt.Println("Not Identical")
  }
}
