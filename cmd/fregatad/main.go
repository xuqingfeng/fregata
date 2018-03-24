package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/influxdata/wlog"
	"github.com/xuqingfeng/fregata/logging"
	"github.com/xuqingfeng/fregata/server"
	"github.com/xuqingfeng/fregata/vars"
)

var (
	version string
	usage   = `usage: ` + vars.DaemonName + ` [flags]
    -version
                    Output version number.
    -config <path>
                    The path of the configuration file.
    `
)

type Main struct {
	Logger *log.Logger
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
		Logger: wlog.New(os.Stdout, "[fregatad] ", log.LstdFlags),
	}
}

// Run accept stdin args, parse flags and start server
func (m *Main) Run(args ...string) error {

	var o options
	flag.BoolVar(&o.version, "version", false, "")
	flag.StringVar(&o.config, "config", "", "")
	flag.Usage = func() {
		fmt.Fprintln(os.Stdout, usage)
	}
	flag.Parse()

	if o.version {
		fmt.Fprintf(os.Stdout, vars.DaemonName+": %s", version+"\n")
		os.Exit(0)
	}
	if o.config != "" {
		// load config from file
		config, err := parseConfig(o.config)
		if err != nil {
			return err
		}
		logService := logging.NewService(config.Logging, os.Stdout, os.Stderr)
		logService.Open()
		defer logService.Close()
		_, err = server.New(config, logService)
		if err != nil {
			return err
		}
	} else {
		flag.Usage()
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
