package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const(
	ERRFORMAT = "sh: %s\n"
)

func main(){

	reader := bufio.NewReader(os.Stdin)

	for{
		fmt.Print("sh> ")
		text, err := reader.ReadString('\n')
		if err!=nil{
			fmt.Println("Error reading text from console: ", err)
		}

		text = strings.ReplaceAll(text, "\n", "")
		args := strings.Split(text, " ")

		executeCommands(args)
	}
}

func executeCommands(args []string) int {
	if len(args) == 0{
		return 1
	}

	for k, v := range builtins {
		if args[0] == k {
			return v(args[1:])
		}
	}

	return launch(args)
}