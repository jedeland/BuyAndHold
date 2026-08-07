package csv

import (
	"fmt"
	"os"
	"regexp"
)

type FileData struct {
	Name    string
	Content []byte
}

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
	var fileArray []FileData

	for _, file := range files {
		name := file.Name()
		if re.MatchString(name) {
			fmt.Println("Relevant File:", name)
		    // Add CSV extraction logic here
			file, err := os.ReadFile(dirPath + "/" + name)
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}
			// Double check
			fileArray = append(fileArray, FileData{Name: name, Content: file})
		}
	}
	// Needs a fixup
	for _, file := range fileArray {
		fmt.Println("File Name:", file.Name)
		fmt.Println("File Content:", string(file.Content))
	}

	normaliseCsvs(fileArray)

}

func normaliseCsvs(fileArray []FileData) {
	// Placeholder for normalisation logic
	// Python code for yahoo uses the following heading 
	// Price,Close,High,Low,Open,Volume, 
	// so the historical data needs to be normalised to this format