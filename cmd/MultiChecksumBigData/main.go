/* MultiChecksum BigData Variant
   MultiChecksum is a tool to get multiple checksums of given files at once.
   This (BigData) variant is implemented in a way it can handle files bigger than your amount of RAM.
*/
package main

import (
	"bufio"
	"bytes"
	"flag"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"golang.org/x/crypto/sha3"
	"fmt"
	"github.com/dchest/blake2b"
	"github.com/dchest/blake2s"
	"io"
	"log"
	"os"
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
	fmt.Printf("Multichecksum CMD (BigData) Version: %s compiled by %s from commit %s at %s\n", version, builtBy, commit, date)
}

// generate and print checksum for each file
// takes a string (filename) as argument
// prints different kinds of checksums for file
//func printSums(filename string) (err error) {
func checksumWorker(w int, jobsChan <-chan string, resultChan chan<- string) {
	for j := range jobsChan {
		log.Printf("started worker %d for '%s'\n", w, j)
		var resultBuf bytes.Buffer
		rw := bufio.NewWriter(&resultBuf)
		// generate handles for all our kinds of checksums
		md5 := md5.New()
		sha1 := sha1.New()
		sha256 := sha256.New()
		sha512 := sha512.New()
		blake2s := blake2s.New256()
		blake2b2 := blake2b.New256()
		blake2b5 := blake2b.New512()
		sha3256 := sha3.New256()
		sha3512 := sha3.New512()
		// create a MultiWriter to write to all handles at once
		w := io.MultiWriter(md5, sha1, sha256, sha512, sha3256, blake2s, blake2b2, blake2b5, sha3512)
		// open file handle
		f, err := os.Open(j)
		if err != nil {
			return
		}
		defer f.Close()
		// copy file to multi writer
		bytesWritten, err := io.Copy(w, f)
		if err != nil {
			return
		}
		// print out checksums
		fmt.Fprintf(rw, "Checksums for %s (Size: %d):\n", j, bytesWritten)
		fmt.Fprintf(rw, "MD5      (%s): %x\n", j, md5.Sum(nil))
		fmt.Fprintf(rw, "SHA1     (%s): %x\n", j, sha1.Sum(nil))
		fmt.Fprintf(rw, "SHA256   (%s): %x\n", j, sha256.Sum(nil))
		fmt.Fprintf(rw, "SHA3-256 (%s): %x\n", j, sha3256.Sum(nil))
		fmt.Fprintf(rw, "Blake2s  (%s): %x\n", j, blake2s.Sum(nil))
		fmt.Fprintf(rw, "Blake2b2 (%s): %x\n", j, blake2b2.Sum(nil))
		fmt.Fprintf(rw, "Blake2b5 (%s): %x\n", j, blake2b5.Sum(nil))
		fmt.Fprintf(rw, "SHA512   (%s): %x\n", j, sha512.Sum(nil))
		fmt.Fprintf(rw, "SHA3-512 (%s): %x\n", j, sha3512.Sum(nil))
		rw.Flush()
		resultChan <- resultBuf.String()
	}
	return
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
	jobsChan := make(chan string, 100)
	resultChan := make(chan string, 100)

	for w := 1; w <= len(args); w++ {
		go checksumWorker(w, jobsChan, resultChan)

	}

	// iterate over arguments given and add filenames to the jobs channel
	for i := 0; i < len(args); i++ {
		filename := args[i]
		// send filename to jobs channel
		jobsChan <- filename
	}
	close(jobsChan)

	// collect results
	for a := 1; a <= len(args); a++ {
		fmt.Printf("%s", <-resultChan)
	}
}
