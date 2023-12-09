package main

import (
	"encoding/csv"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("arquivo-de-origem.csv")
	checkErr(err)

	reader := csv.NewReader(file)
	labels, err := reader.Read()
	checkErr(err)
	records, err := reader.ReadAll()
	checkErr(err)

	sortNames(records)
	records = append([][]string{labels}, records...)

	writeOutputFile(records)
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

func writeOutputFile(records [][]string) {
	file, err := os.Create("arquivo-destino.csv")
	checkErr(err)

	writer := csv.NewWriter(file)
	writer.WriteAll(records)
}
