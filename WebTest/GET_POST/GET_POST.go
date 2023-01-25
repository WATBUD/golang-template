package GET_POST

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadT1Stock(filepath string, url string, jsonString string) error {

	resp, err := http.Post("https://www.ufs-portal.com/api/T1StockData/getT1StockCasePricekg", "application/json", bytes.NewReader([]byte(jsonString)))
	if err != nil {
		fmt.Println(string("handle error"))
	}

	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println(string("io error"))
	// }
	// fmt.Println(string(body))
}
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
func HttpGet() {

	resp, err := http.Get("https://tw.yahoo.com/")
	if err != nil {
		fmt.Println(string("handle error"))
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string("io error"))
	}

	fmt.Println(string(body))
}
