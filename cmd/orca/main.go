package main

import (
	"fmt"
	"os"

	"github.com/defendops/orca/pkg/cmd/factory"
	"github.com/defendops/orca/pkg/cmd/root"
)

var updaterEnabled = ""

type exitCode int

const (
	exitOK      exitCode = 0
	exitError   exitCode = 1
	exitCancel  exitCode = 2
	exitAuth    exitCode = 4
	exitPending exitCode = 8
)

func executeORCA() exitCode {
	cmdFactory := factory.NewCmdFactory("ORCA Beta")

	if err := root.ExecuteRootCmd(cmdFactory); err != nil {
		fmt.Println("Error executing root command:", err)
		return exitError
	}
	return exitOK
}

func main() {
	code := executeORCA()
	os.Exit(int(code))
}
