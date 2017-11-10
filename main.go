package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
	"os"
)

const usageTemplate = `
NAME:
   tf-release-info - Write terraform plugin release-info JSON

USAGE:
   tf-release-info PLUGIN-VERSION TERRAFORM-VERSION JSON-FILE-PATH

ARGS:
   PLUGIN-VERSION    : terraform plugin version via "v1.0.0"
   TERRAFORM-VERSION : terraform version via "v1.0.0"
   JSON-FILE-PATH    : file path of target JSON

COPYRIGHT:
   Copyright (C) 2017 Kazumichi Yamamoto.
`

func main() {

	if err := validateArgs(os.Args); err != nil {
		printError(err)
		usage()
		os.Exit(1)
	}

	pluginVersion := os.Args[1]
	terraformVersion := os.Args[2]
	filePath := os.Args[3]

	f, err := os.Open(filePath)
	if err != nil {
		printError(err)
		os.Exit(2)
	}

	releaseInfo, err := readReleaseInfo(f)
	if err != nil {
		printError(err)
		os.Exit(2)
	}

	info := newReleaseInfo(terraformVersion, pluginVersion)
	f.Close()

	for _, r := range releaseInfo {
		if r.PluginVersion == info.PluginVersion {
			printWarn(fmt.Sprintf("Plugin %q is already exists.", info.PluginVersion))
			os.Exit(0)
		}
	}
	releaseInfo = append(releaseInfo, info)

	f, err = os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		printError(err)
		os.Exit(3)
	}
	defer f.Close()

	if err := writeReleaseInfo(f, releaseInfo); err != nil {
		printError(err)
		os.Exit(3)
	}

	printInfo(fmt.Sprintf("Added Plugin %q info to JSON", info.PluginVersion))
	os.Exit(0)
}

func validateArgs(args []string) error {
	if len(args) != 4 {
		return fmt.Errorf("Invalid arguments num")
	}
	// file exists?
	filePath := args[3]
	if _, err := os.Stat(filePath); err != nil {
		return fmt.Errorf("File(%q) is not exists", filePath)
	}

	return nil
}

func printWarn(msg string) {
	stderr := colorable.NewColorableStderr()
	clr := color.New(color.FgHiMagenta)
	clr.Fprintf(stderr, "[WARN] %s\n", msg)
}
func printInfo(msg string) {
	stderr := colorable.NewColorableStderr()
	clr := color.New(color.FgHiBlue)
	clr.Fprintf(stderr, "[INFO] %s\n", msg)
}
func printError(err error) {
	stderr := colorable.NewColorableStderr()
	clr := color.New(color.FgHiRed)
	clr.Fprintf(stderr, "[ERROR] %s\n", err)
}

func usage() {
	fmt.Printf(usageTemplate)
}
