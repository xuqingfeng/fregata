package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/influxdata/wlog"
	"github.com/xuqingfeng/fregata/logging"
	"github.com/xuqingfeng/fregata/server"
	"github.com/xuqingfeng/fregata/vars"
)

var (
	usage = `usage: ` + vars.DaemonName + ` [flags]
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
	version bool
	config  string
}

func main() {

	m := NewMain()
	if err := m.Run(os.Args[1:]...); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

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
	flag.BoolVar(&o.version, "version", false, "")
	flag.StringVar(&o.config, "config", "", "")
	flag.Usage = func() {
		fmt.Fprintln(os.Stdout, usage)
	}
	flag.Parse()

	if o.version {
		fmt.Fprintf(os.Stdout, vars.DaemonName+": %s", vars.Version+"\n")
		os.Exit(0)
	}
	if o.config != "" {
		// load config from file
		config, err := parseConfig(o.config)
		if err != nil {
			return err
		}
		logService := logging.NewService(config.Logging, m.Stdout, m.Stderr)
		logService.Open()
		defer logService.Close()
		_, err = server.New(config, logService)
		if err != nil {
			return err
		}
	}

	return nil
}

func parseConfig(path string) (*server.Config, error) {

	if path == "" {
		return nil, errors.New("No configuration provided")
	}

	config := server.NewConfig()
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}

	return config, nil
}
