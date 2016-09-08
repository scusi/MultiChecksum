MultiChecksum
=============

MultiChecksum is a go module and a commandline application to generate 
multiple cryptographic checksums at once for a given set of files.

Library Variant
===============

There is a very simple library called `multichecksum`. It consists out of one 
function and one type.

```
// Checksums is a map that holds filename and checksums for that file
type Checksums map[string]string

// CalcChecksums takes a filename and the file content and returns a map with the checksums
func CalcChecksums(filename string, data []byte) *Checksums {
    [...] }
```

The [command line tool](cmd/MultiChecksum/MultiChecksum.go) makes use of that 
library as you can see in the following listing:

```
        [...]
		chksums := multichecksum.CalcChecksums(filename, data)
		for _, sum := range *chksums {
			fmt.Printf("%s", sum)
		}
        [...]
```

Stand Alone Variant
===================

The stand alone variant is under `cmd/MultiChecksumNoLib/`. This is a
command line tool that does the job, no library required.

This is basically the first version of MultiChecksum

Fetch
=====

```
go get github.com/scusi/MultiChecksum
```

Install
=======

```
# Install the lib
cd ${GOPATH}/src/github.com/scusi/MultiChecksum
go install

# Install the commandline app
cd ${GOPATH}/src/github.com/scusi/MultiChecksum/cmd/MultiChecksum/
go install MultiChecksum.go
```

Run
===

After installing the commandline app you should be able to execute

```
MultiChecksum file1.bin file2.bin anotherFile.doc
```


