package http

import "net/http"

func RegisterRouter(userHandler UserHandler) {
	http.HandleFunc("/user", userHandler.GetUser)
}
