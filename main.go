package main

import (
	"fmt"
	"time"
	"net/http"
)

func main() {
	now := time.Now().Unix()
	client := &http.Client{}

	url := fmt.Sprintf("https://query1.finance.yahoo.com/v1/finance/download/%%5EDJI?period1=694224000&period2=%d&interval=1d&events=history", now)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response Status:", resp.Status)
}