package main

import (
	"net/http"
	"github.com/brendanHwang/golang/web/3강/myapp"
)

func main() {

	http.ListenAndServe(":8080", myapp.NewHttpHandler()) // 8080 포트를 듣고 있는 서버 실행
}
