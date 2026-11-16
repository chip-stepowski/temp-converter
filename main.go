package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func roundToOne(val float64) float64 {
	return math.Round(val*10) / 10
}

func main() {

	// creates a reader with methods
	reader := bufio.NewReader(os.Stdin)

	// add loop so the user can continue entering conversions until q is pressed
	for {
		fmt.Print("Enter a temperature and a unit of measurement (c, f, or k) or 'q' to quit: ")

		// read the users input until a newline (Enter). it keeps the newline. _ is used to omit the error value, as we don't need it
		line, _ := reader.ReadString('\n')

		// trim the whitespace from both ends, the newline, and then we lowercase it
		input := strings.ToLower(strings.TrimSpace(line))

		// check if input is q. if it is, exit the loop
		if input == "q" {
			fmt.Println("See you later!")
			break
		}

		// check the input length
		if len(input) == 0 {
			fmt.Println("Please enter some value.")
			continue
		}

		// take the clean user input and create a slice from index 5 to the end (index 5 is f or c)
		unit := input[len(input)-1:]
		// create a slice from clean user input, starting at the beginning up until 5 (the f or c), then trim the ending whitespace
		numberPart := strings.TrimSpace(input[:len(input)-1])

		// convert numberPart string into float64
		i, err := strconv.ParseFloat(numberPart, 64)

		// make sure numberPart is a number
		if err != nil {
			fmt.Println("Please enter a valid value.")
			continue
		}

		// convert F to C with switch
		switch unit {
		case "f":
			c := (i - 32) * 5 / 9
			fmt.Println(i, "degrees fahrenheit is equal to", roundToOne(c), "degrees celsisus.")
		case "c":
			f := i*9/5 + 32
			fmt.Println(i, "degrees celsius is equal to", roundToOne(f), "degrees fahrenheit.")
		case "k":
			f := (i-273.15)*9/5 + 32
			c := i - 273.15
			fmt.Println(i, "degrees Kelvin is equal to", roundToOne(f), "degrees fahrenheit and ", roundToOne(c), "degress celsius.")

		default:
			fmt.Println("Please enter a valid unit of measurement.")
		}
	}

}
