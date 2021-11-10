package main

import (
	"os"

	"github.com/atlantis0/androidthreatbeat/cmd"

	_ "github.com/atlantis0/androidthreatbeat/include"
)

func setDefaultNS2(addrs []string) {
	defaultNS = addrs
}

func main() {

	// set default DNS,
	// workaround since android dosen't have /etc/resolve.conf file
	// https://github.com/coyove/goflyway/issues/126
	setDefaultNS2([]string{"8.8.8.8:53", "1.1.1.1:53"})

	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
