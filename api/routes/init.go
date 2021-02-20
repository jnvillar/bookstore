package routes

import "net/http"

func registerRoutes(){
	http.HandleFunc("/ping", heartbeat)
	http.ListenAndServe(":8080", nil)
}