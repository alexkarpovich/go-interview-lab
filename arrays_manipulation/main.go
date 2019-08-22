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
	beg int32 = iota
	end
	val
)

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {
	var maxVal, tmp int64
	var id int32
	arr := make([]int64, n+1)

	for _, q := range queries {
		arr[q[beg]] = arr[q[beg]] + int64(q[val])

		id = q[end] + 1
		if id <= n {
			arr[id] = arr[id] - int64(q[val])
		}
	}

	tmp = 0
	for _, v := range arr {
		tmp = tmp + v

		if maxVal < tmp {
			maxVal = tmp
		}
	}

	return maxVal
}

func main() {
	f, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReaderSize(f, 1024*1024)

	stdout := os.Stdout
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != int(3) {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

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
