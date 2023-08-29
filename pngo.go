package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type PngChunk struct {
	length uint32
	ctype  uint32
	data   []byte
	crc    uint32
}

func main() {
	dat, err := os.ReadFile("img.png")
	check(err)

	if binary.BigEndian.Uint64(dat[0:8])&uint64(0x89504e470d0a1a0a) != 0x89504e470d0a1a0a {
		panic("PNG header not valid!")
	}

	var cstart uint32 = 8
	for {
		var length uint32 = binary.BigEndian.Uint32(dat[cstart : cstart+4])
		var ctype string = string(dat[cstart+4 : cstart+8])
		var data []byte = dat[cstart+8 : cstart+8+length]
		var crc []byte = dat[cstart+8+length : cstart+12+length]

		fmt.Printf("TYPE: %s\n", ctype)
		fmt.Printf("DATA: %x\n", data)
		fmt.Printf("CRC: %x\n\n", crc)

		if ctype == "IEND" {
			break
		}

		cstart += 12 + length
	}

}
