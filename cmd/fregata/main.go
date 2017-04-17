package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xuqingfeng/fregata/vars"
)

var (
	usage = `usage: ` + vars.Name + `[flags]
	-version
			Output version number.
	`

	version bool
)

func main() {

	flag.BoolVar(&version, "version", false, "")
	flag.Usage = func() {
		fmt.Fprintln(os.Stdout, usage)
	}
	flag.Parse()

	if version {
		fmt.Fprintf(os.Stdout, vars.Name+": %s", vars.Version+"\n")
		os.Exit(0)
	}

	// TODO: extend
	/*
			- list enabled services
		        - show status
	*/
}
