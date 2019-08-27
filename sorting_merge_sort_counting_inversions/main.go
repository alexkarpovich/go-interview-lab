package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var count int64

func merge(arr []int32, lid, mid, rid int) {
	i := lid
	j := mid + 1
	k := 0
	res := make([]int32, rid-lid+1)

	for ok := true; ok; ok = i <= mid || j <= rid {
		if j > rid && i <= mid || i <= mid && arr[i] <= arr[j] {
			res[k] = arr[i]
			i++
		} else if j <= rid {
			res[k] = arr[j]

			if i <= mid && j <= rid && arr[j] < arr[i] {
				count += int64(mid - i + 1)
			}

			j++
		}
		k++
	}

	copy(arr[lid:], res)
}

func mergeSort(arr []int32, lid, rid int) {
	if lid < rid {
		mid := (rid - lid) / 2
		mergeSort(arr, lid, lid+mid)
		mergeSort(arr, lid+mid+1, rid)
		merge(arr, lid, lid+mid, rid)
	}
}

// Complete the countInversions function below.
func countInversions(arr []int32) int64 {
	count = 0
	mergeSort(arr, 0, len(arr)-1)

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

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
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

		result := countInversions(arr)

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
