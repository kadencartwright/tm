package main

import (
	"fmt"
	"os"

	"tm/cmd"
)

func main() {
	if err := cmd.NewRootCmd(cmd.DefaultDependencies()).Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
