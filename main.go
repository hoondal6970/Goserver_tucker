package main

import (
	"net/http"

	"github.com/hoondal6970/dockertest/myapp"
)

func main() {

	http.ListenAndServe(":1234", myapp.NewHttpHandler()) //서버를 실행하는 구문. 자동으로 listen 상태로 실행되고 있는 상태이다.

}
