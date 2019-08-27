package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the isValid function below.
func isValid(s string) string {
	var c1, c2 int
	count := make(map[rune]int)
	freq := make(map[int]int)

	for _, c := range s {
		count[c]++
	}

	for _, cnt := range count {
		freq[cnt]++

		if c1 == 0 {
			c1 = cnt
		}

		if c2 == 0 && c1 != cnt {
			c2 = cnt
		}
	}

	if len(freq) < 2 || len(freq) == 2 && (freq[c1] == 1 && freq[c2] > 1 && c2 > c1 || freq[c1] > 1 && freq[c2] == 1 && c2 < c1) {
		return "YES"
	}

	return "NO"
}

func main() {
	f, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReaderSize(f, 1024*1024)
	stdout := os.Stdout
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	result := isValid(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
