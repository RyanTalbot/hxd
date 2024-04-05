package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
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
		err = binary.Read(buffer, binary.BigEndian, &out)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", hex.Dump(out))
	}
}
