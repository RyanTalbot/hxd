package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		os.Exit(1)
	}

	littlEndianFlag := flag.Bool("e", false, "set for little Endian")

	flag.Parse()

	var byteOrder binary.ByteOrder
	if *littlEndianFlag {
		byteOrder = binary.LittleEndian
	} else {
		byteOrder = binary.BigEndian
	}

	file, err := os.Open(args[len(args)-1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileBytes := make([]byte, 256)
	out := make([]byte, 256)

	for {
		_, err := file.Read(fileBytes)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		buffer := bytes.NewBuffer(fileBytes)
		err = binary.Read(buffer, byteOrder, &out)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", hex.Dump(out))
	}
}
