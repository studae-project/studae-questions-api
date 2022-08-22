package controller

import "net/http"

type Controller interface {
	POST(writer http.ResponseWriter, request *http.Request)
}
