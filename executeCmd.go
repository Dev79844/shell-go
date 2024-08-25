package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

)

func Execute(args []string) int{
	exit :=  strings.Compare(args[0], "exit"); if exit==0 {
		os.Exit(0)
	}

	return executeSimpleCommand(args)
}

func executeSimpleCommand(args []string) int {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = nil
	
	if err := cmd.Run(); err!=nil{
		fmt.Println(args)
		fmt.Printf(ERRFORMAT, err.Error())
	}
	return 1
}