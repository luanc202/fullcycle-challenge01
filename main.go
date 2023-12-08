package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
	}

	reader := csv.NewReader(file)
	labels, err := reader.Read()
	if err != nil {
		fmt.Println(err)
	}
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i][0] < records[j][0]
	})

	records = append([][]string{labels}, records...)

	file, err = os.Create("arquivo-destino.csv")
	if err != nil {
		fmt.Println(err)
	}

	writer := csv.NewWriter(file)
	writer.WriteAll(records)
}
