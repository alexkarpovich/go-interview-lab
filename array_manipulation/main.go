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
	BEG int = iota
	END
	VAL
)

type Point struct {
	X       int32
	Related int32
	Value   int64
}

func Init(q []int32, s int) *Point {
	p := new(Point)
	p.Value = int64(q[VAL])

	if s == BEG {
		p.X = q[BEG]
		p.Related = q[END]
	} else {
		p.X = q[END]
		p.Related = q[BEG]
	}

	return p
}

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {
	var maxVal int64
	var p, t *Point
	var i int32
	seg := []int{BEG, END}
	rng := make([]*Point, n+1)

	for qid, q := range queries {
		if qid == 0 {
			for _, s := range seg {
				p = Init(q, s)
				rng[q[s]] = p
			}

			maxVal = p.Value
		} else {
			for _, s := range seg {
				if rng[q[s]] != nil {
					p = rng[q[s]]
					p.Value = p.Value + int64(q[VAL])
				} else {
					p := Init(q, s)

					if s == BEG {
						for i = q[BEG] - 1; i > 0; i-- {
							t = rng[i]

							if t != nil && t.Related > q[BEG] {
								p.Value = p.Value + t.Value
								break
							}
						}
					} else {
						for i = q[END] + 1; i <= n; i++ {
							t = rng[i]

							if t != nil && t.X > q[END] && t.Related < q[END] {
								p.Value = p.Value + t.Value
								break
							}
						}
					}

					rng[q[s]] = p
				}

				if p.Value > maxVal {
					maxVal = p.Value
				}
			}

			for i = q[BEG] + 1; i < q[END]; i++ {
				t = rng[i]

				if t != nil {
					t.Value = t.Value + int64(q[VAL])

					if t.Value > maxVal {
						maxVal = t.Value
					}
				}
			}
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
