package userInput

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

//ReadInput This function reads the users input, and prints
func ReadInput() {

	read := bufio.NewReader(os.Stdin)
	fmt.Print("Please input a letter:")
	inp, _ := read.ReadString('\n')
	fmt.Print(inp)

}
