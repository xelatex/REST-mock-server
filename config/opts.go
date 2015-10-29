// Copyright 2015-2018 Arthur Chunqi Li. All rights reserved.

package config

import (
	"fmt"
	"os"
)

type Options struct {
	Host        string `json:"addr"`
	Port        int    `json:"port"`
	ControlPort int    `json:"control_port"`
	ConfigPath  string `json:"config_path"`
}

var usageStr = `
Server Options:
    -a, --addr HOST                  Bind to HOST address (default: 0.0.0.0)
    -p, --port PORT                  Use PORT for clients (default: 1080)
    -P, --control_port PORT          Use PORT for clients to control (default: 1070)
    -c, --configPath PATH            Use PATH to hold configuration files (default: ./config)

Common Options:
    -h, --help                       Show this message
    -v, --version                    Show version
`

// Usage will print out the flag options for the server.
func Usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}
