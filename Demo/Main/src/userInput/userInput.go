package userInput

import (
  "fmt"

	"bufio"
	"os"

)

//ReadInput This function reads the users input, and prints

func ReadInput() string {

	read := bufio.NewReader(os.Stdin)
	fmt.Print("Please input a letter:")
	inp, _ := read.ReadString('\n')
	inputFMT := checkString(inp)


	return inputFMT

}

func checkString(inp string) string {
	inputFMT := inp[0:len(inp)-1]
	return inputFMT
}
