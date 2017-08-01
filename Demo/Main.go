package main 



func main() {

	read := bufio.NewReader(os.Stdin)
    fmt.Print("Please input a letter:")
    inp, _ := read.ReadString('\n')
    fmt.Println(inp)

	
	

}