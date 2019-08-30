package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	down int32 = -1
	up   int32 = 1
)

// Complete the substrCount function below.
func substrCount(n int32, s string) int64 {
	var step, swing int32
	count := int64(n)

	for i := int32(0); i < n-1; i++ {
		step = 1
		swing = up

		for j := i + 1; j < n; j++ {
			if s[j] == s[i] {
				step += swing

				if swing == up && step > 1 || swing == down && step == 0 {
					count++
				}
			}

			if s[j] != s[i] {
				if swing == down {
					break
				}
				swing = down
			}
		}
	}

	return count
}

func main() {
	f, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReaderSize(f, 1024*1024)
	stdout := os.Stdout
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	s := readLine(reader)

	result := substrCount(n, s)

	fmt.Fprintf(writer, "%d\n", result)

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
