package utils

import (
	"bufio"
	"os"
)

func WriteToFile(fname string, lines []string) {
	f, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buffer := bufio.NewWriter(f)
	for _, line := range lines {
		_, err := buffer.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}
	if err := buffer.Flush(); err != nil {
		panic(err)
	}
}
