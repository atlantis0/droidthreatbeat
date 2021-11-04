package main

import (
	"os"

	"github.com/atlantis0/androidthreatbeat/cmd"

	_ "github.com/atlantis0/androidthreatbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
