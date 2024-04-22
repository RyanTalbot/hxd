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
	"regexp"
	"strings"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		os.Exit(1)
	}

	littlEndianFlag := flag.Bool("e", false, "set for little Endian")
	groupSizeFlag := flag.Int("g", 4, "separate the output of every x bytes by a whitespace")

	flag.Parse()

	var byteOrder binary.ByteOrder
	groupingSize := *groupSizeFlag
	if *littlEndianFlag {
		byteOrder = binary.LittleEndian
		groupingSize = *groupSizeFlag
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
		fmt.Printf("%s", getLine(out, groupingSize*2))
	}
}

func getLine(data []byte, groupingSize int) string {
	dump := hex.Dump(data)

	regex := `^([0-9A-Fa-f]+)\s+((?:[0-9A-Fa-f]{2}\s*){1,16})\s*\|\s*(.+?)\s*\|$`
	re := regexp.MustCompile(regex)

	lines := strings.Split(dump, "\n")

	var result strings.Builder
	for _, line := range lines {
		if line == "" {
			continue
		}
		matches := re.FindStringSubmatch(line)
		if matches != nil && len(matches) > 3 {
			offset := matches[1]
			hexData := matches[2]
			text := matches[3]

			hexData = strings.Replace(hexData, " ", "", -1)
			hexData = insertSpaces(hexData, groupingSize)

			fmt.Fprintf(&result, "%s: %s   %s\n", offset, hexData, text)
		}
	}
	return result.String()
}

func insertSpaces(hexData string, groupSize int) string {
	if groupSize < 1 {
		return hexData
	}

	var builder strings.Builder
	for i := 0; i < len(hexData); i += groupSize {
		if i+groupSize < len(hexData) {
			builder.WriteString(hexData[i:i+groupSize] + " ")
		} else {
			builder.WriteString(hexData[i:])
		}
	}
	return builder.String()
}
