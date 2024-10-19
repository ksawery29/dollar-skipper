package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func runCommand(args []string) {
	var cmd *exec.Cmd

	// set the command to use based on the OS
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", strings.Join(args, " "))
	} else {
		cmd = exec.Command("/bin/sh", "-c", strings.Join(args, " "))
	}

	// set the stdout, stderr and stdin to the current process
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	cmd.Env = os.Environ()
	err := cmd.Run()

	if err != nil {
		fmt.Printf("Error running command: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	// get the args
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "dollar-skipper: No command provided")
		os.Exit(1)
	}

	// and run the command
	runCommand(args)
}
