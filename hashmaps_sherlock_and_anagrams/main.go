package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func countCharPairs(n int) int32 {
	return int32(n * (n - 1) / 2)
}

func countInRestSubstr(beg int, s string, ss string) int32 {
	var count int32
	var checksum int32
	lnss := len(ss)
	ln := len(s) - lnss

	for i := beg; i <= ln; i++ {
		checksum = 0
		for _, c := range ss {
			checksum = checksum + int32(c)*int32(c)*(c-'a'+1)*(c-'a'+1)
		}

		for _, c := range s[i : i+lnss] {
			checksum = checksum - int32(c)*int32(c)*(c-'a'+1)*(c-'a'+1)
		}

		if checksum == 0 {
			count++
		}
	}

	return count
}

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	var count int32
	chars := make(map[rune][]int)

	for i, c := range s {
		chars[c] = append(chars[c], i)
	}

	for _, ids := range chars {
		if len(ids) > 1 {
			count = count + countCharPairs(len(ids))
		}
	}

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			count = count + countInRestSubstr(i+1, s, s[i:j+1])
		}
	}

	return count
}

func main() {
	f, err := os.Open("./test_.txt")
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
