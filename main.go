package main

import (
	"log"
	"os/exec"

	"github.com/juanpickselov/go-test-exec/funshell"
)

func main() {
	stdout, err := funshell.ShellCommand(exec.Command)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(stdout.String())
}
