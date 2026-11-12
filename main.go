package main

import (
	"fmt"
	"math"
)

func roundToOne(val float64) float64 {
	return math.Round(val*10) / math.Round(10)
}

func main() {
	// we use float64 as using an int would result in 9 / 5 as 1
	var i float64
	var unit string

	fmt.Print("Enter a temperature and a unit of measurement (C or F) or 'q' to quit: ")
	// Scan reads input, each variables type decides when it stops reading. First, it tries to match an integer (i). When it gets to unit, it stops and fills that variable with values that can be a string
	fmt.Scan(&i, &unit)

	// convert F to C with switch
	switch unit {
	case "f":
		c := (i - 32) * 5 / 9
		fmt.Println(i, "degrees fahrenheit is equal to", roundToOne(c), "degrees celsisus.")
	case "c":
		f := i*9/5 + 32
		fmt.Println(i, "degrees celsius is equal to", roundToOne(f), "degrees fahrenheit.")
	default:
		fmt.Println("Please enter a valid unit of measurement.")
	}

}

// ask the user for a number and a unit (C or F) // done
// convert and print the result with one decimal place // done
// reject unknown units // done
// keep converting until the user types 'q'
// extra: add Kelvin
// DoD: when it can handle bad input without crashing
