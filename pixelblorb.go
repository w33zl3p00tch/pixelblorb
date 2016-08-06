/*
pixelblorb v0.0.1

A tool to extract pixel data from images


Homepage: https://github.com/w33zl3p00tch/pixelblorb


Copyright (c) 2016 Manuel Iwansky ( w33zl3p00tch [at) gmail d0t com )
released under a BSD-style license
*/

package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strconv"
)

func main() {
	// flags
	imgFilePtr := flag.String("f", "", "name of the image file to open")
	outFmtPtr := flag.String("format", "d",
		"output format; d for decimal, b for binary;\n"+
			"order: [X, Y, R, G, B, A]")

	flag.Parse()

	imgFile := *imgFilePtr
	outFmt := *outFmtPtr

	// open input file
	inFile, err := os.Open(imgFile)
	check(err)
	defer inFile.Close()

	img, _, err := image.Decode(inFile)
	check(err)

	// if -d is set, parse the image with decimal output and exit
	if outFmt == "d" {
		imgParseDec(img)
	} else if outFmt == "b" { // output in binary, Big Endian order
		imgParseBin(img)
	}
}

// imgParseBin outputs pixel data in binary representation.
// The order is Big Endian.
func imgParseBin(src image.Image) {
	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			curPixel := src.At(x, y)

			srcr, srcg, srcb, alph := curPixel.RGBA()
			r := uint8(srcr)
			g := uint8(srcg)
			b := uint8(srcb)
			a := uint8(alph)

			rBin := strconv.FormatInt(int64(r), 2)
			gBin := strconv.FormatInt(int64(g), 2)
			bBin := strconv.FormatInt(int64(b), 2)
			aBin := strconv.FormatInt(int64(a), 2)

			rBin = binPad(rBin)
			gBin = binPad(gBin)
			bBin = binPad(bBin)
			aBin = binPad(aBin)

			fmt.Println(x, y, rBin, gBin, bBin, aBin)
		}
	}
}

// binPad pads the byte with zeros until it consists of 8 bits, if necessary.
func binPad(binVal string) string {
	for len(binVal) < 8 {
		binVal = "0" + binVal
	}

	return binVal
}

// imgParseDec outputs pixel data as decimal values.
func imgParseDec(src image.Image) {
	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			curPixel := src.At(x, y)

			srcr, srcg, srcb, alph := curPixel.RGBA()
			r := uint8(srcr / 256)
			g := uint8(srcg / 256)
			b := uint8(srcb / 256)
			a := uint8(alph / 256)

			fmt.Println(x, y, r, g, b, a)
		}
	}
}

// check for errors and quit if an error occured.
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
