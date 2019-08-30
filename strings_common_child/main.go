package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// Complete the commonChild function below.
func commonChild(s1 string, s2 string) int32 {
	var count, commonCount int32
	var prevk, tlen int
	var ch byte
	ln1 := len(s1)
	ln2 := len(s2)
	chars := make(map[byte][]int)

	for i, c := range s1 {
		ch = byte(c)
		chars[ch] = append(chars[ch], i)
	}

	for i := 0; i < ln2; i++ {
		count = 0
		prevk = 0

		for j := i; j < ln2; j++ {
			ids, ok := chars[s2[j]]
			tlen = len(ids)

			if ok {
				k := sort.Search(tlen, func(ix int) bool { return ids[ix] > prevk })

				if k >= 0 && k < tlen {
					prevk = ids[k]
					count++

					if ids[k] == ln1-1 {
						break
					}
				}
			} else if j == i {
				continue
			}
		}

		if count > commonCount {
			commonCount = count
		}
	}

	return commonCount
}

func main() {
	f, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReaderSize(f, 1024*1024)
	stdout := os.Stdout
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s1 := readLine(reader)

	s2 := readLine(reader)

	result := commonChild(s1, s2)

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
