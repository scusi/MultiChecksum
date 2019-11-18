package main

import (
	"fmt"
	"github.com/scusi/MultiChecksum"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// get command line arguments (without our own name)
	args := os.Args[1:]
	// print how many files we where given
	fmt.Println("Number of Files given: ", len(args))
	// iterate over arguments given and call printSums for each filename
	for i := 0; i < len(args); i++ {
		filename := args[i]
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		sum := multichecksum.CalcChecksums(filename, data)
		fmt.Printf("Checksums for %s:\n", sum.Filename)
		for _, h := range sum.Hashes {
			fmt.Printf(" %s %s%s\t%x\n", sum.Filename, h.HashName, Spaces(7-len(h.HashName)), h.Hash)
		}

	}

}

// Spaces produces a given number of space
func Spaces(i int) (s string) {
	for c := 0; c < i; c++ {
		s = s + " "
	}
	return
}
