package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	var i, j int32
	var minSwaps int32 = 0
	var ln int32 = int32(len(arr))

	pos := make(map[int32]int32)

	for i = 0; i < ln; i++ {
		pos[arr[i]-1] = i
	}

	for i = 0; i < ln; i++ {
		if arr[i] != i+1 {
			j, _ = pos[i]
			arr[i], arr[j] = arr[j], arr[i]
			pos[arr[j]-1] = j
			minSwaps++
		}
	}

	return minSwaps
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

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := minimumSwaps(arr)

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
