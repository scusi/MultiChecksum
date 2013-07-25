MultiChecksum
=============

a tool implemented in go, to show multiple checksums for given file(s).
Currently MultiChecksum can generate MD5, SHA1, SHA2 and SHA5.

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

Install
=======

<pre>
 cd MultiChecksum
 go install
</pre>
Usage
=====

The tool takes files as arguments and prints MD5, SHA1, SHA2 and SHA5, like shown in the example below.

<pre>
nyx:MultiChecksum flow$ MultiChecksum README.md 
Number of Files given:  1
Checksums for README.md:
MD5    (README.md): 0936beea77781d50a6d691624efb1278
SHA1   (README.md): fda7a80f2b7420b50a319a42c77213f3b5502ba2
SHA256 (README.md): 741a117a6d7cbec6cb1fb9f7f9816ff8d3569a9de7a81cab24dd73f4a4e52cb9
SHA512 (README.md): 38fa4bb1f7548aea3e892daebba888a55d5e0950f29ee92c80b6f222551ad150865a29e55490b894258f1f713d57e12d93f04eb17ef3a784560d6fea78e4da1f
</pre>

Note: MultiChecksum allows to be given more than one filename as commandline argument.
