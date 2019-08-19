package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func countCombinations(n int) int32 {
	return int32(n * (n - 1) / 2)
}

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	var count int32
	chars := make(map[rune][]int)

	for i, c := range s {
		chars[c] = append(chars[c], i)
	}

	for i, c := range s {
		ids, _ := chars[c]

		if len(ids) > 1 {
			count = count + countCombinations(len(ids))
		}
	}

}

func main() {
	f, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReaderSize(f, 1024*1024)
	stdout := os.Stdout
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

		fmt.Fprintf(writer, "%d\n", result)
	}

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
