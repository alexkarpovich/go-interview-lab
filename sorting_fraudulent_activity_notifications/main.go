package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func median(arr []int32) float64 {
	mid := float64(len(arr)-1) / 2
	midl := int(mid)
	midr := int(math.Ceil(mid))

	if midl != midr {
		return float64(arr[midl]+arr[midr]) / 2
	}

	return float64(arr[midl])
}

func insert(s []int32, at int, val int32) []int32 {
	if at >= len(s) {
		return append(s, val)
	}
	// Make sure there is enough room
	s = append(s, 0)
	// Move all elements of s up one slot
	copy(s[at+1:], s[at:])
	// Insert the new element at the now free position
	s[at] = val
	return s
}

func remove(slice []int32, s int) []int32 {
	return append(slice[:s], slice[s+1:]...)
}

// Complete the activityNotifications function below.
func activityNotifications(expenditure []int32, d int32) int32 {
	var count int32
	i := d
	wnd := make([]int32, d)
	ln := int32(len(expenditure))

	copy(wnd, expenditure[0:d])
	sort.Slice(wnd, func(i, j int) bool { return wnd[i] < wnd[j] })

	for ok := true; ok; ok = i < ln-1 {
		if float64(expenditure[i]) >= 2*median(wnd) {
			count++
		}

		rem := sort.Search(len(wnd), func(j int) bool { return wnd[j] >= expenditure[i-d] })
		wnd = remove(wnd, rem)
		ins := sort.Search(len(wnd), func(j int) bool { return wnd[j] >= expenditure[i] })
		wnd = insert(wnd, ins, expenditure[i])
		i++
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

	nd := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	expenditureTemp := strings.Split(readLine(reader), " ")

	var expenditure []int32

	for i := 0; i < int(n); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int32(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}

	result := activityNotifications(expenditure, d)

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
