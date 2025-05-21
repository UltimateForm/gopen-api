package api

import "net/http"

func characterHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello"))
}
