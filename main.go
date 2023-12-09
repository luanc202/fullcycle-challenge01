package main

import (
	"encoding/csv"
	"os"
	"sort"
)

func main() {
	args := os.Args[1:]
	file, err := os.Open(args[0])
	checkErr(err)

	reader := csv.NewReader(file)
	labels, err := reader.Read()
	checkErr(err)
	records, err := reader.ReadAll()
	checkErr(err)

	sortNames(records)
	records = append([][]string{labels}, records...)

	writeOutputFile(records, args[1])
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func sortNames(records [][]string) {
	sort.Slice(records, func(i, j int) bool {
		return records[i][0] < records[j][0]
	})
}

func writeOutputFile(records [][]string, filename string) {
	file, err := os.Create(filename)
	checkErr(err)

	writer := csv.NewWriter(file)
	writer.WriteAll(records)
}
