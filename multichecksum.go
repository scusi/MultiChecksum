// Package multichecksum does calculate a bunch of checksums for a given data at once.
// The at once calulation of the checksums is realized useing an io.MultiWriter.
//
// For an example on how to use this package please have a look at cmd/MultiChecksum/MultiChecksum.go.
// If you want to have the same functionality in a single go source file, please look at:
// cmd/MultiChecksumNoLib/MultiChecksum.go
//
// Supported checksums are: 
//   - MD5, 
//   - SHA-1, 
//   - SHA-2, 
//   - SHA-3 (32 and 64 byte),
//   - SHA-5, 
//   - Blake2s, 
//   - Blake2b (32 and 64 byte)
//
package multichecksum // import "github.com/scusi/MultiChecksum"

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"golang.org/x/crypto/sha3"
	"github.com/dchest/blake2b"
	"github.com/dchest/blake2s"
	"io"
)

// MultiChecksum object to store all hashes for a given file
type MultiChecksum struct {
	Filename string
	Hashes   []Hashsum
}

// Hashsum object to store a given hashing algorithm and a hash.
type Hashsum struct {
	HashName string
	Hash     []byte
}

// CalcChecksums takes a filename and the file content and returns a map with the checksums
func CalcChecksums(filename string, data []byte) *MultiChecksum {
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
	w := io.MultiWriter(md5, sha1, sha256, sha512, sha3256, sha3512, blake2s, blake2b2, blake2b5)
	// write (file) content to our MultiWriter (w)
	w.Write(data)
	// create a map and write filename and checksums to it
	msc := &MultiChecksum{
		Filename: filename,
		Hashes: []Hashsum{
			Hashsum{HashName: "MD5",
				Hash: md5.Sum(nil)},
			{HashName: "SHA1",
				Hash: sha1.Sum(nil)},
			{HashName: "SHA256",
				Hash: sha256.Sum(nil)},
			{HashName: "SHA3-256",
				Hash: sha3256.Sum(nil)},
			{HashName: "Blake2s",
				Hash: blake2s.Sum(nil)},
			{HashName: "Blake2b",
				Hash: blake2b2.Sum(nil)},
			{HashName: "SHA512",
				Hash: sha512.Sum(nil)},
			{HashName: "SHA3-512",
				Hash: sha3512.Sum(nil)},
			{HashName: "Blake5",
				Hash: blake2b5.Sum(nil)},
		},
	}
	return msc
}
