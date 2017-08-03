package readInput

import (
	"fmt"
	"bufio"
	"os"
)

func readInput() {

	read := bufio.NewReader(os.Stdin)
    fmt.Print("Please input a letter:")
    inp, _ := read.ReadString('\n')
    fmt.Print(inp)

}