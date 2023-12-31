package main

import (
	"net/http"

	"brendan.com/myproject/myapp"
)

func main() {

	http.ListenAndServe(":8000", myapp.NewHttpHandler()) // 8080 포트를 듣고 있는 서버 실행
}
