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
		chksums := multichecksum.CalcChecksums(filename, data)
		for _, sum := range *chksums {
			fmt.Printf("%s", sum)
		}
	}

}
