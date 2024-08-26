package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

		args, ok := parseLine(text)
		if ok && args!=nil{
			executeCommands(args)
		}
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

func parseLine(line string) ([]string, bool){
	args := regexp.MustCompile("'(.+)'|\"(.+)\"|\\S+").FindAllString(line, -1)
	for i, arg := range args {
		if (arg[0] == '"' && arg[len(arg)-1] == '"') || (arg[0] == '\'' && arg[len(arg)-1] == '\'') {
			args[i] = arg[1 : len(arg)-1]
		}
	}
	
	if args[0] == "export"{
		if len(args) == 1{
			fmt.Printf(ERRFORMAT, "arguments need for export")
		}

		exportArgs := strings.Split(args[1], "=")
		if len(exportArgs) != 2 {
			fmt.Printf(ERRFORMAT, "wrong format of export")
			return nil, false
		}
		os.Setenv(exportArgs[0], exportArgs[1])
		return args, false
	}

	fmt.Println(args)

	for i, arg := range args {
		if arg[0] == '$' {
			args[i] = os.Getenv(arg[1:])
		}
	}

	return args, true
}