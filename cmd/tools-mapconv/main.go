package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/opd-ai/BrowserQuest/pkg/gamemap"
)

func main() {
	input := flag.String("in", "", "path to a BrowserQuest JSON or world_client.js map")
	output := flag.String("out", "", "path to the converted Go map JSON")
	pretty := flag.Bool("pretty", true, "write indented JSON")
	flag.Parse()

	if *input == "" || *output == "" {
		fmt.Fprintln(os.Stderr, "usage: tools-mapconv -in <map-file> -out <output-file> [-pretty=true]")
		os.Exit(2)
	}

	raw, err := os.ReadFile(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read input: %v\n", err)
		os.Exit(1)
	}

	world, err := gamemap.Parse(raw)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse input: %v\n", err)
		os.Exit(1)
	}

	payload, err := world.Marshal(*pretty)
	if err != nil {
		fmt.Fprintf(os.Stderr, "marshal output: %v\n", err)
		os.Exit(1)
	}

	if err := os.MkdirAll(filepath.Dir(*output), 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "create output directory: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(*output, payload, 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "write output: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("converted %s -> %s\n", *input, *output)
}
