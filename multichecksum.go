// Package multichecksum does calculate a bunch of checksums for a given data at once.
// The at once calulation of the checksums is realized useing an io.MultiWriter.
//
// For an example on how to use this package please have a look at cmd/MultiChecksum/MultiChecksum.go.
// If you want to have the same functionality in a single go source file, please look at:
// cmd/MultiChecksumNoLib/MultiChecksum.go
//
// Supported checksums are: MD5, SHA1, SHA2, SHA5, Blake2s, Blake2b (32 byte) and Blake2b (64 byte)
//
package multichecksum

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"github.com/dchest/blake2b"
	"github.com/dchest/blake2s"
	"io"
)

// Checksums is a map that holds filename and checksums for that file
type Checksums map[string]string

// CalcChecksums takes a filename and the file content and returns a map with the checksums
func CalcChecksums(filename string, data []byte) *Checksums {
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
	w.Write(data)
	// create a map and write filename and checksums to it
	sums := make(map[string]string)
	sums["Filename"] = fmt.Sprintf("Checksums for '%s':\n", filename)
	sums["MD5"] = fmt.Sprintf("MD5      (%s): %x\n", filename, md5.Sum(nil))
	sums["SHA1"] = fmt.Sprintf("SHA1     (%s): %x\n", filename, sha1.Sum(nil))
	sums["SHA2"] = fmt.Sprintf("SHA256   (%s): %x\n", filename, sha256.Sum(nil))
	sums["Blake2s"] = fmt.Sprintf("Blake2s  (%s): %x\n", filename, blake2s.Sum(nil))
	sums["Blake2b"] = fmt.Sprintf("Blake2b2 (%s): %x\n", filename, blake2b2.Sum(nil))
	sums["Blake2b5"] = fmt.Sprintf("Blake2b5 (%s): %x\n", filename, blake2b5.Sum(nil))
	sums["SHA512"] = fmt.Sprintf("SHA512   (%s): %x\n", filename, sha512.Sum(nil))
	// type conversion - convert our map to our Checksums datatype
	chksums := Checksums(sums)
	// return a Checksums datatype map with the result sums
	return &chksums
}

func (cs *Checksums) String() string {
	var outbuf bytes.Buffer
	w := bufio.NewWriter(&outbuf)
	for typ, sum := range *cs {
		if typ == "Filename" {
			continue
		}
		fmt.Fprint(w, "%s", sum)
	}
	w.Flush()
	return outbuf.String()
}

func (cs *Checksums) Filter(types ...string) (string, error) {
	for _, typ := range types {
		for ctyp, sum := range *cs {
			if ctyp == typ {
				return sum, nil
			}
		}
	}
	err := fmt.Errorf("type not found")
	return "", err
}
