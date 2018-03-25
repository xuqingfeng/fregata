package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xuqingfeng/fregata/vars"
)

var (
	version = "master"
	usage   = `usage: ` + vars.Name + `[flags]
	-version
			Output version number.
	`

	versionFlag bool
)

func main() {

	flag.BoolVar(&versionFlag, "version", false, "")
	flag.Usage = func() {
		fmt.Fprintln(os.Stdout, usage)
	}
	flag.Parse()

	if versionFlag {
		fmt.Fprintf(os.Stdout, vars.Name+": %s", version+"\n")
		os.Exit(0)
	}

	// TODO: extend
	/*
			- list enabled services
		        - show status
	*/
}
