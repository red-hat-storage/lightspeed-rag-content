package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/red-hat-storage/lightspeed-rag-content/internal/download"
	"github.com/red-hat-storage/lightspeed-rag-content/internal/parse"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  go run ./cmd download --version 4.22 [--force]")
		fmt.Println("  go run ./cmd parse    --version 4.22")
		os.Exit(1)
	}

	subcommand := os.Args[1]

	switch subcommand {
	case "download":
		fs := flag.NewFlagSet("download", flag.ExitOnError)
		version := fs.String("version", "", "ODF documentation version")
		force := fs.Bool("force", false, "Redownload existing files")
		fs.Parse(os.Args[2:])

		if *version == "" {
			fmt.Println("Usage: go run ./cmd download --version 4.22 [--force]")
			os.Exit(1)
		}

		if err := download.Run(*version, *force); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

	case "parse":
		fs := flag.NewFlagSet("parse", flag.ExitOnError)
		version := fs.String("version", "", "ODF documentation version")
		fs.Parse(os.Args[2:])

		if *version == "" {
			fmt.Println("Usage: go run ./cmd parse --version 4.22")
			os.Exit(1)
		}

		if err := parse.Run(*version); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

	default:
		fmt.Printf("Unknown command: %s\n", subcommand)
		fmt.Println("Available commands: download, parse")
		os.Exit(1)
	}
}
