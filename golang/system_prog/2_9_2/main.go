package main

import (
	// "compress/gzip"

	"encoding/csv"
	"log"
	"os"
)

func main() {

	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	w := csv.NewWriter(os.Stdout)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	w1 := csv.NewWriter(file)
	for _, record := range records {
		if err := w1.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	w1.Flush()
}
