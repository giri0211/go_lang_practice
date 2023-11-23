package web_service

import (
	"io"
	"net/http"
	"os"
)

// Handler reads the data from file and write to http Response.
func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./web-service/menu.txt")
	io.Copy(w, f)
}

// holding the control here for the we server to keep listening
func Webserver_start() {
	http.HandleFunc("/", Handler) //localhost:3000 path
	// http.HandleFunc("/test", Handler) //localhost:3000/test path
	http.ListenAndServe("localhost:3000", nil)
}
