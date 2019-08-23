package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
	var count int64
	possibleTuples := make(map[int64]int64)
	possibleTriples := make(map[int64]int64)

	for _, a := range arr {
		if _, ok := possibleTriples[a]; ok {
			// We met 3rd item and possibleTriples turnes into real
			count += possibleTriples[a]
		}

		if _, ok := possibleTuples[a]; ok {
			// Imagine we meet a*r as next then we'll have possibleTriples[a*r] triples
			possibleTriples[a*r] += possibleTuples[a]
		}
		// Imagine we meet a*r as next then we'll have possibleTuples[a*r] tuples
		possibleTuples[a*r]++
	}

	return count
}

func main() {
	f, err := os.Open("./test_.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReaderSize(f, 16*1024*1024)
	stdout := os.Stdout
	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(nr[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	r, err := strconv.ParseInt(nr[1], 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItem)
	}

	ans := countTriplets(arr, r)

	fmt.Fprintf(writer, "%d\n", ans)

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
