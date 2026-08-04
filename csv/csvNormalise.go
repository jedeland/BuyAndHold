package csv

import (
	"fmt"
	"os"
	"regexp"
)

func TargetCsvs() {
	// Assumes CSVs have already been unzipped using csvSources
	// Non-normalised CSVs currently 08.04.2026 are
	// NIKKEI225.csv and DJA-Timestamp-20260802.csv

	dirPath := "sources_unzipped/source_data"
	re := regexp.MustCompile(`^[^_]+$`)

	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	fileArray := make([]byte, 0)

	for _, file := range files {
		name := file.Name()
		if re.MatchString(name) {
			fmt.Println("Relevant File:", name)
		    // Add your CSV extraction logic here
			file, err := os.ReadFile(dirPath + "/" + name)
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}
			// Double check
			fileArray = append(fileArray, file...)
		}
	}
	// Needs a fixup
	for fileName, fileContent := range fileArray {
		fmt.Println("File Name:", fileName)
		fmt.Println("File Content:", string(fileContent))
	}


}