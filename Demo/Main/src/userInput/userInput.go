package userInput

import (
  "fmt"

	"bufio"
	"os"
	"convert"
)

//ReadInput This function reads the users input, and prints

func ReadInput() {

	read := bufio.NewReader(os.Stdin)
	fmt.Print("Please input a letter:")
	inp, _ := read.ReadString('\n')
	inputFMT := checkString(inp)
	fmt.Println(inputFMT)
	convert.Convert()
	if inputFMT == "s" {
		fmt.Println("correctumundo")
	} else {
		fmt.Println("wrong")
	}

}

func checkString(inp string) string {
	inputFMT := inp[0:len(inp)-1]
	return inputFMT
}
