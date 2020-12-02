package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFileToArr(file_path string) []string {
	array := make([]string, 0)

	fi, err := os.Open(file_path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return array
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		array = append(array, string(a))
	}
	return array
}

func WriteArrToFile(arr []string, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, value := range arr {
		fmt.Fprintln(w, value)
	}
	return w.Flush()
}
