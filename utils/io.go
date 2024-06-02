package utils

import (
	"bufio"
	"os"
	"path/filepath"
)

func WriteToFile(fname string, text string) {
	fpath := filepath.Dir(fname)
	if _, err := os.Stat(fpath); os.IsNotExist(err) {
		os.MkdirAll(fpath, os.ModePerm)
	}
	f, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buffer := bufio.NewWriter(f)
	if _, err := buffer.WriteString(text); err != nil {
		panic(err)
	}
	if err := buffer.Flush(); err != nil {
		panic(err)
	}
}
