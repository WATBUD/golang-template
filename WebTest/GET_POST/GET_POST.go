package GET_POST

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func HttpPostWithJson() {

	jsonString := `{
  "from": "2022-12-31T16:00:00.000Z",
  "to": "2022-12-31T16:00:00.000Z"
}`

	http.Post("https://www.ufs-portal.com/api/T1StockData/getT1StockCasePricekg", "application/json", bytes.NewReader([]byte(jsonString)))

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
