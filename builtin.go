package main

import (
	"fmt"
	"os"
	"strings"
)

var builtins = map[string]func([]string)int{
	"exit": exit,
	"cd": cd,
}

func exit(args []string) int {
	fmt.Println("Exiting shell...")
	os.Exit(0)
	return 1
}

func cd(args []string) int {
	if len(args) == 0 {
		fmt.Printf(ERRFORMAT, "Enter a path to change directory")
	}else if len(args) > 1 {
		fmt.Printf(ERRFORMAT, "Too many arguments")
	}else{
		err := os.Chdir(args[0])
		if err != nil {
			fmt.Printf(ERRFORMAT, err.Error())
			return 0
		}

		wd, _ := os.Getwd()
		wdSlice := strings.Split(wd, "/")
		os.Setenv("CWD", wdSlice[len(wdSlice)-1])
	}

	return 1
}