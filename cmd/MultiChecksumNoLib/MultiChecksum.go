// MultiChecksum is an implementation in go to show multiple checksums for given files
//
// Author: fw@snurn.de
//
package main

// Import needed packages
import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"golang.org/x/crypto/sha3"
	"fmt"
	"github.com/dchest/blake2b"
	"github.com/dchest/blake2s"
	"github.com/zeebo/blake3"
	"io"
	"io/ioutil"
	"os"
	"flag"
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
	fmt.Printf("Multichecksum CMD (NoLib) Version: %s compiled by %s from commit %s at %s\n", version, builtBy, commit, date)
}

// loader - returns the content of a given file as []byte
// takes a string (the filename of the file to read) as argument
// returns a []byte (content of file) and error
func loader(filename string) (content []byte, err error) {
	// TODO: Reading content entierly into memory does not work in cases
	// - where the content is larger than the available memory
	// - when there is a restriction on maximum available memory for the process
	// This could be overcome by streaming the content into the MultiWriter in func PrintSums
	content, err = ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// generate and print checksum for each file
// takes a string (filename) as argument
// prints different kinds of checksums for file
func printSums(filename string) {
	// call 'loader' to load the file and return it's content as []byte
	content, err := loader(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	// generate handles for all our kinds of checksums
	md5 := md5.New()
	sha1 := sha1.New()
	sha256 := sha256.New()
	sha512 := sha512.New()
	blake2s := blake2s.New256()
	blake2b2 := blake2b.New256()
	blake2b5 := blake2b.New512()
	b3 := blake3.New()
	sha3256 := sha3.New256()
	sha3512 := sha3.New512()
	// create a MultiWriter to write to all handles at once
	w := io.MultiWriter(md5, sha1, sha256, sha512, sha3256, sha3512, blake2s, blake2b2, blake2b5, b3)
	// write (file) content to our MultiWriter (w)
	w.Write(content)
	// print out checksums
	fmt.Printf("Checksums for %s:\n", filename)
	fmt.Printf("MD5        (%s): %x\n", filename, md5.Sum(nil))
	fmt.Printf("SHA1       (%s): %x\n", filename, sha1.Sum(nil))
	fmt.Printf("SHA256     (%s): %x\n", filename, sha256.Sum(nil))
	fmt.Printf("SHA3-256   (%s): %x\n", filename, sha3256.Sum(nil))
	fmt.Printf("Blake2s    (%s): %x\n", filename, blake2s.Sum(nil))
	fmt.Printf("Blake2b2   (%s): %x\n", filename, blake2b2.Sum(nil))
	fmt.Printf("Blake3-256 (%s): %x\n", filename, b3.Sum(nil))
	fmt.Printf("Blake2b5   (%s): %x\n", filename, blake2b5.Sum(nil))
	fmt.Printf("SHA512     (%s): %x\n", filename, sha512.Sum(nil))
	fmt.Printf("SHA3-512   (%s): %x\n", filename, sha3512.Sum(nil))
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
		fmt.Printf("flags: %v", args)
		// print how many files we where given
		fmt.Println("Number of Files given: ", len(args))
	}
	// iterate over arguments given and call printSums for each filename
	for i := 0; i < len(args); i++ {
		filename := args[i]
		printSums(filename)
	}
}
