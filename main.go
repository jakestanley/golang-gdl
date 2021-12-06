package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	ini "gopkg.in/ini.v1"
)

var (
	commands map[string]func()
	supports []string
)

type args struct {
	port string
}

func main() {

	initCommands()
	initSupports()

	osCfgDir, err := os.UserConfigDir()
	cfgDir := fmt.Sprintf("%s/jakestanley.github.io/gdl", osCfgDir)
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll(cfgDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	cfgPath := fmt.Sprintf("%s/config.ini", cfgDir)

	fmt.Printf("Using config file: '%s'\n", cfgPath)

	cfg, err := ini.LooseLoad(cfgPath)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	section, err := cfg.GetSection("Awesome")
	if err != nil {
		section, _ = cfg.NewSection("Awesome")
	}

	section.NewKey("Hello", "Balls!")

	err = cfg.SaveTo(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	var command string
	var sourcePort string

	exe := os.Args[0]

	fmt.Printf("Usage: '%s'\n", exe)

	command = os.Args[1]

	actionSet := false

	for k, _ := range commands {
		if k == command {
			fmt.Printf("Command: '%s' is supported \n", k)
			actionSet = true
			break
		}
	}

	// kinda hacking flags
	if actionSet {
		os.Args = os.Args[1:]
	}

	defaultPort := supports[0]
	flag.StringVar(&sourcePort, "source-port", defaultPort, "the source port you wish to use")
	flag.Parse()

	fmt.Println(sourcePort)
	flag.Usage()
}

func initCommands() {

	commands = make(map[string]func())

	configure := func() {}

	commands["configure"] = configure
}

func initSupports() {

	/*
	   there are "families" with similar syntax. might need to use that,
	   i.e the GL variant of prboom, or ZDoom and LZDoom
	*/
	supports = append(supports, "gzdoom", "prboom")
}
