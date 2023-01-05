package main

import (
	customHttp "WebTest/CustomHttp"
	"fmt"
)

func main() {
	fmt.Println("Hello world") //這個寫入到 w 的是輸出到客戶端的

	customHttp.StartListenServe()
}
