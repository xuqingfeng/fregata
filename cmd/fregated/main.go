package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/influxdata/wlog"
	"github.com/xuqingfeng/fregata/vars"
)

var (
	usage = `usage: ` + vars.DaemonName + ` run [flags]
    -version
                    Output version number.
    -config <path>
                    The path of the configuration file.
    `
)

type Main struct {
	Logger *log.Logger
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

type options struct {
	version string
	config  string
}

func main() {

	m := NewMain()

}

func NewMain() *Main {

	return &Main{
		Logger: wlog.New(os.Stderr, "[fregated] ", log.LstdFlags),
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

func (m *Main) Run(args ...string) error {

	var o options
	f := flag.NewFlagSet("", flag.ContinueOnError)
	f.StringVar(&o.version, "version", false, "")
	f.StringVar(&o.config, "config", "", "")
	f.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
	}
	if err := f.Parse(args); err != nil {

		if o.version {
			fmt.Printf(vars.DaemonName+": %s", vars.Version)
			os.Exit(0)
		}

	}

}
