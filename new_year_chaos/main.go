package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Max(x, y int32) int32 {
	if x < y {
		return y
	}
	return x
}

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	var min_bribes int32 = 0
	var i int32
	var j int32
	ln := int32(len(q))

	for i = ln - 1; i >= 0; i-- {

		if q[i]-i > 3 {
			fmt.Println("Too chaotic")
			return
		}

		for j = Max(0, q[i]-2); j < i; j++ {
			if q[j] > q[i] {
				min_bribes++
			}
		}
	}

	fmt.Println(min_bribes)
}

func main() {
	f, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReaderSize(f, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
