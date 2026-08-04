package main

import (
	"BuyAndHold/csv"
)

func main() {
	csv.UnzipCSV("sources/source_data.tar.gz")
	csv.TargetCsvs()
}