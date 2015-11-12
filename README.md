MultiChecksum
=============

a tool implemented in go, to show multiple checksums for given file(s).
Currently MultiChecksum can generate MD5, SHA1, SHA2, SHA5 and Blake2.

MultiChecksum runs on the following operating systems and platforms.
Operating Systems supported: Windows, Linux, FreeBSD, NetBSD, OpenBSD, Mac OS X
Platforms supported: i386, amd64, arm 

Requirements
============

You need to have go installed on your machine. For more informtion about go check out http://www.golang.org/

Download
========

You can clone the source using git
<pre>
 git clone https://github.com/scusi/MultiChecksum.git
</pre>

or you can download the source as a ZIP Archive from:

 https://github.com/scusi/MultiChecksum/archive/master.zip

and unpack it into any local directory

 unzip -e master.zip

Install
=======

After download you change into the new MultiChecksum directory and call _go install_

<pre>
 cd MultiChecksum
 go install
</pre>
Usage
=====

The tool takes files as arguments and prints MD5, SHA1, SHA2, SHA5, Blake2s and Blake2b checksums 
like shown in the example below.

<pre>
nyx:MultiChecksum flow$ MultiChecksum doc.go 
Number of Files given:  1
Checksums for doc.go:
MD5     (doc.go): 4773d77cc5299500ea5c3c9c0201bc4c
SHA1    (doc.go): 894466465df48d3fba5508cc8480a12abfed6920
SHA256  (doc.go): 6a8b25c0c9195c5fddf74ee7bad56306082c9027fae366bc73801f71e2e150e5
SHA512  (doc.go): fa18a57cac11b9c816c7c44b99787255d16253c6bb1b08f6154386e929b8ef7f1d7156fc6699ed7fd9bc4a4dbf95421597a790d34837948651761ec72dfca4ad
Blake2s (doc.go): 03371343e73bf58702388ae9e6aa9c38678b95fab8140169553f70f52a110165
Blake2b (doc.go): 774a991d33cf75e05a5b9536d15557586ecb3c116da0a16f1f5a352f175ee44d
</pre>

Note: MultiChecksum allows to be given more than one filename as commandline argument.

Webapp
======

There is also a Webapp variant for docker of _MultiChecksum_ available from https://github.com/scusi/MultiChecksumWeb
