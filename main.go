// Copyright 2015-2018 Arthur Chunqi Li. All rights reserved.

package main

import (
	"flag"
	"log"
	"os"

	"./config"
	"./server"
)

func main() {
	var showVersion bool
	// Server options
	opts := config.Options{}

	// Parse flags
	flag.IntVar(&opts.Port, "port", 1080, "Port to listen on.")
	flag.IntVar(&opts.Port, "p", 1080, "Port to listen on.")
	flag.IntVar(&opts.ControlPort, "controlPort", 1070, "Port to listen on.")
	flag.IntVar(&opts.ControlPort, "P", 1070, "Port to listen on.")
	flag.StringVar(&opts.Host, "addr", "0.0.0.0", "Network host to listen on.")
	flag.StringVar(&opts.Host, "a", "0.0.0.0", "Network host to listen on.")
	flag.BoolVar(&showVersion, "v", false, "Print version information.")
	flag.StringVar(&opts.ConfigPath, "c", "./config", "Configuration path to hold data.")
	flag.StringVar(&opts.ConfigPath, "configPath", "./config", "Configuration path to hold data.")

	flag.Usage = config.Usage

	flag.Parse()

	// Show version and exit
	if showVersion {
		server.PrintServerAndExit()
	}

	if opts.Port == opts.ControlPort {
		log.Printf("ERROR: port and control port (%d) cannot be the same.\n", opts.Port)
		os.Exit(-1)
	}

	if _, err := os.Stat(opts.ConfigPath); os.IsNotExist(err) {
		log.Fatalf("no such file or directory: %s\n", opts.ConfigPath)
		return
	}

	log.Printf("host: %s\n", opts.Host)
	log.Printf("port: %d\n", opts.Port)
	log.Printf("control port: %d\n", opts.ControlPort)

	// Create the server with appropriate options.
	s := server.New(opts)
	// Start things up. Block here until done.
	s.Start()
}
