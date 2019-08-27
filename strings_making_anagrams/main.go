package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the makeAnagram function below.
func makeAnagram(a string, b string) int32 {
	var commonChars int32
	aCounter := make(map[rune]int)

	for _, c := range a {
		aCounter[c]++
	}

	for _, c := range b {
		cnt, ok := aCounter[c]

		if ok && cnt > 1 {
			aCounter[c]--
			commonChars++
		} else if ok {
			delete(aCounter, c)
			commonChars++
		}
	}

	return int32(len(a)+len(b)) - commonChars*2
}

func main() {
	f, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReaderSize(f, 1024*1024)
	stdout := os.Stdout
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	a := readLine(reader)

	b := readLine(reader)

	res := makeAnagram(a, b)

	fmt.Fprintf(writer, "%d\n", res)

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
