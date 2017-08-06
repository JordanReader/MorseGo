package main

import (
	"userInput"
	"convert"
	"fmt"
)

func main() {

	userLetter := userInput.ReadInput()
	convertedLetter := convert.Convert(userLetter)
	fmt.Println(userLetter + " in morse is " + convertedLetter)

}




