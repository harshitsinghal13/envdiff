package cmd

import (
	"flag"
	"fmt"
	"os"
)

var registry []subcommand

type subcommand struct {
	name        string
	description string
	run         func(args []string)
}

func register(sc subcommand) {
	registry = append(registry, sc)
}

func Execute() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	name := os.Args[1]

	if name == "--help" || name == "-h" {
		printHelp()
		os.Exit(0)
	}

	for _, sc := range registry {
		if sc.name == name {
			sc.run(os.Args[2:])
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown command: %q\n\n", name)
	printHelp()
	os.Exit(1)
}

func printHelp() {
	fmt.Println("envdiff — compare .env files and catch config drift")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  envdiff <command> [flags]")
	fmt.Println()
	fmt.Println("Available commands:")
	for _, sc := range registry {
		fmt.Printf("  %-12s %s\n", sc.name, sc.description)
	}
	fmt.Println()
	fmt.Println("Use envdiff <command> --help for more info.")
}

func newFlags(name string, args []string) *flag.FlagSet {
	fs := flag.NewFlagSet(name, flag.ExitOnError)
	return fs
}