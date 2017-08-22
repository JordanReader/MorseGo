package userInput

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

//ReadInputE This function reads the users english input, and prints the converted morse phrase.

func ReadInputE() string {

	read := bufio.NewReader(os.Stdin)
	fmt.Print("Please input a letter:")
	inp, _ := read.ReadString('\n')
	inputFMT := checkString(inp)

	return strings.ToLower(inputFMT)

}

//ReadInputM This function reads the users morse input, and prints the converted english phrase.

func ReadInputM() string {

	read := bufio.NewReader(os.Stdin)
	fmt.Print("Please input a morse letter:")
	inp, _ := read.ReadString('\n')
	inputFMT := checkString(inp)




	return inputFMT

}

func checkString(inp string) string {
	inputFMT := inp[0:len(inp)-1]
	return inputFMT
}
