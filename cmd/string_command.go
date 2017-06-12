package cmd

import (
	"fmt"
)

type stringCommand struct {
	command string
}

func (sc *stringCommand) execute() error {
	fmt.Printf(sc.command)
	return nil
}

func (sc *stringCommand) parse(arguments []string) error {
	return nil
}
