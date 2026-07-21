package main

import (
	"fmt"
	"github.com/scusi/MultiChecksum"
	"os"
	"flag"
	"log"
)

var (
	version	string
	commit	string
	date	string
	builtBy	string
)

var showVersion bool
var beVerbose	bool

func init() {
	flag.BoolVar(&showVersion, "version", false, "shows version info, and exits")
	flag.BoolVar(&beVerbose, "verbose", false, "be verbose")
}

func VersionInfo() {
	fmt.Printf("Multichecksum CMD Version: %s compiled by %s from commit %s at %s\n", version, builtBy, commit, date)
}

func main() {
	flag.Parse()
	if showVersion {
		VersionInfo()
		os.Exit(0)
	}
	// get command line arguments (without our own name)
	args := flag.Args()
	if beVerbose {
		log.Printf("flags: %v", args)
		// print how many files we where given
		fmt.Println("Number of Files given: ", len(args))
	}
	// iterate over arguments given and call printSums for each filename
	hasError := false
	for i := 0; i < len(args); i++ {
		filename := args[i]
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", filename, err)
			hasError = true
			continue
		}
		sum, err := multichecksum.CalcChecksums(filename, data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calculating checksums for %s: %v\n", filename, err)
			hasError = true
			continue
		}
		fmt.Printf("Checksums for %s:\n", sum.Filename)
		for _, h := range sum.Hashes {
			fmt.Printf(" %s %s%s\t%x\n", sum.Filename, h.HashName, Spaces(7-len(h.HashName)), h.Hash)
		}
	}
	if hasError {
		os.Exit(1)
	}
}

// Spaces produces a given number of space
func Spaces(i int) (s string) {
	for c := 0; c < i; c++ {
		s = s + " "
	}
	return
}
