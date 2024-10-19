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

	// loop through args and add quotes around arguments that contain spaces
	for i, arg := range args {
		// if the argument contains a space and isnt already wrapped in quotes, add them
		if strings.Contains(arg, " ") && !strings.HasPrefix(arg, "\"") && !strings.HasSuffix(arg, "\"") {
			args[i] = "\"" + arg + "\""
		}
	}

	// set the command to use based on the OS
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", append([]string{"/C"}, args...)...)
	} else {
		cmd = exec.Command("/bin/sh", append([]string{"-c"}, args...)...)
	}

	// set the stdout, stderr and stdin to the current process
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	cmd.Env = os.Environ()
	err := cmd.Run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "dollar-skipper: Error running command: %v\n", err)
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
