package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

)

func Execute(cmd string){
	exit :=  strings.Compare(cmd, "exit"); if exit==0 {
		os.Exit(0)
	}

	output, err := exec.Command(cmd).Output()
	if err!=nil{
		fmt.Printf(ERRFORMAT, err)
	}
	fmt.Fprint(os.Stdout, string(output[:]))
	fmt.Print("\n")
}