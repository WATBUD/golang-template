package customHttp

import (
	"fmt"
	"net/http"
)

func Test(w http.ResponseWriter, _request *http.Request) {
	fmt.Fprintf(w, "Enter Test") //這個寫入到 w 的是輸出到客戶端的
}

func SayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Enter SayhelloName") //這個寫入到 w 的是輸出到客戶端的
}

func Testlog(_rs http.ResponseWriter, _request *http.Request) {

	fmt.Fprintf(_rs, "Enter Testlog")

}

func Testlog2(_rs http.ResponseWriter, _request *http.Request) {

	fmt.Fprintf(_rs, "Enter Testlog2")

}
func ErrorPage(_rs http.ResponseWriter, _request *http.Request) {

	fmt.Fprintf(_rs, "Enter ErrorPage")

}
func StartListenServe() {
	fmt.Println("Enter StartListenServe")
	http.HandleFunc("/", SayhelloName) //設定存取的路由
	http.HandleFunc("/Test", Test)     //設定存取的路由
	http.HandleFunc("/Test", Testlog)
	http.HandleFunc("/Test2", Testlog2)
	http.HandleFunc("/", ErrorPage)
	http.ListenAndServe(":9090", nil) //設定監聽的埠
}
