package main

import (
	"userInput"
	"convert"
	"flag"
	"fmt"
)

func main() {

	parseFlag()



}
func parseFlag() {
	lang := flag.String("lang", "", "a string")
	flag.Parse()

	if *lang == "e" {
		English()
	} else if *lang == "m" {
		Morse()
	} else {
		parseFlag()
	}
}

func English() {
	input := userInput.ReadInputE()
	convertedInput := convert.ConvertE(input)
	fmt.Println(input + " in morse is " + convertedInput)
}

func Morse() {
	input := userInput.ReadInputM()
	convertedInput := convert.ConvertM(input)
	fmt.Println(input + " english is " + convertedInput)
}







