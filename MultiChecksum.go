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
	"fmt"
	"github.com/dchest/blake2b"
	"github.com/dchest/blake2s"
	"io"
	"io/ioutil"
	"os"
)

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
	// create a MultiWriter to write to all handles at once
	w := io.MultiWriter(md5, sha1, sha256, sha512, blake2s, blake2b2, blake2b5)
	// write (file) content to our MultiWriter (w)
	w.Write(content)
	// print out checksums
	fmt.Printf("Checksums for %s:\n", filename)
	fmt.Printf("MD5      (%s): %x\n", filename, md5.Sum(nil))
	fmt.Printf("SHA1     (%s): %x\n", filename, sha1.Sum(nil))
	fmt.Printf("SHA256   (%s): %x\n", filename, sha256.Sum(nil))
	fmt.Printf("Blake2s  (%s): %x\n", filename, blake2s.Sum(nil))
	fmt.Printf("Blake2b2 (%s): %x\n", filename, blake2b2.Sum(nil))
	fmt.Printf("Blake2b5 (%s): %x\n", filename, blake2b5.Sum(nil))
	fmt.Printf("SHA512   (%s): %x\n", filename, sha512.Sum(nil))
}

func main() {
	// get command line arguments (without our own name)
	args := os.Args[1:]
	// print how many files we where given
	fmt.Println("Number of Files given: ", len(args))
	// iterate over arguments given and call printSums for each filename
	for i := 0; i < len(args); i++ {
		filename := args[i]
		printSums(filename)
	}
}
