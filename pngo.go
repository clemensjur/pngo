package main

import (
	"encoding/binary"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("img.png")
	check(err)

	if binary.BigEndian.Uint64(dat[0:8])&uint64(0x89504e470d0a1a0a) != 0x89504e470d0a1a0a {
		panic("PNG header not valid!")
	}
}
