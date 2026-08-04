package csv

import (
	"archive/tar"
	"compress/gzip"
	"path/filepath"
	"fmt"
	"io"
	"os"
)

func UnzipCSV(tarPath string) {
	tarFile, err := os.Open(tarPath)
	if err != nil {
		fmt.Println("Error opening tar file:", err)
		return
	}
	defer tarFile.Close()

	gzipReader, err := gzip.NewReader(tarFile)
	if err != nil {
		fmt.Println("Error opening tar file:", err)
		return
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		head, err := tarReader.Next()
		// Error at end of tar file, break the loop
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading tar entry:", err)
			return
		}

		unpackTarget := filepath.Join("sources_unzipped", head.Name)

		if head.Typeflag == tar.TypeDir {
			if err := os.MkdirAll(unpackTarget, 0755); err != nil {
				fmt.Println("Error creating directory:", err)
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(unpackTarget), 0755); err != nil {
			fmt.Println("Error creating parent directory:", err)
			continue
		}

		tarContent, err := io.ReadAll(tarReader)
		if err != nil {
			fmt.Println("Error reading file content:", err)
			return
		}

		fmt.Println("File:", head.Name)
		if err := os.WriteFile(unpackTarget, tarContent, 0644); err != nil {
			fmt.Println("Error writing file:", err)
		}
	}
}