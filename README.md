# pepperpickle


## Synopsis

PIXELBLORB is a command line tool to extract pixel information from image files. Although its intended purpose is to aid in steganalysis, it can be used to convert image files to text files for further processing.
True to the Unix philosophy of doing one thing and doing it well, it simply outputs a list of data that can be piped into other external commands.

Supported input formats are JPEG, PNG and GIF. The output can be either decimal or in binary represantation and is a space-separated list of x, y, color and alpha values:

```0 0 234 143 4 255```

```0 1 233 142 4 255```

```0 2 233 140 6 255```

```... and so on. The Format is [X Y R G B A].```


## Usage

to output an image file's pixel data as decimal values:

```$ pixelblorb -f IMAGE.jpg```



to output an image's pixel data in binary representation:

```$ pixelblorb -format b -f IMAGE.png```



to print a short help:

```$ pixelblorb -h```



## Installation

Binaries for Windows, Linux-amd64, Mac OS X and FreeBSD-amd64 are available at:
https://github.com/w33zl3p00tch/pixelblorb/releases

Simply extract the binary to a folder in your PATH, e.g. /usr/local and make sure that it is executable.



## Installation from source

pixelblorb is written in Go 1.6.

```$ go build pixelblorb.go```

or install with

```$ go install pixelblorb```



## Revision history
 ```
v0.0.1: initial commit
```


## License

pixelblorb is released under a BSD-Style license.
