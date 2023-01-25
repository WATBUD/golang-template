package main

import (
	//customHttp "WebTest/CustomHttp"
	GET_POST "WebTest/GET_POST"
	"fmt"
)

//var fileUrl string

func main() {
	//customHttp.StartListenServe()
	//GET_POST.HttpGet()
	//GET_POST.HttpPostWithJson()

	fileUrl := "https://www.ufs-portal.com/api/T1StockData/getT1StockCasePricekg"
	//fileUrl := "https://gophercoding.com/img/logo-original.png"
	jsonString := `{
    "from": "2022-11-01T16:00:00.000Z",
    "to": "2022-12-31T16:00:00.000Z"
    }`

	err := GET_POST.DownloadT1Stock("T1Stock.xlsx", fileUrl, jsonString)
	if err != nil {
		fmt.Println("Error downloading file: ", err)
		return
	}

	fmt.Println("Downloaded: " + fileUrl)
}
