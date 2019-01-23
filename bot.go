package main

import (
	"./payload"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func jsonHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		fmt.Printf("method is not POST")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Header.Get("Content-Type") != "application/json" {
		fmt.Printf("Content-Type is not application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		fmt.Printf("Content-Length error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := make([]byte, length)
	length, err = req.Body.Read(body)
	if err != nil && err != io.EOF {
		fmt.Printf("body error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var jsonBody payload.PullRequestPayload
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		fmt.Printf("jsonBody error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("%v\n", jsonBody)

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/json", jsonHandler)
	http.ListenAndServe(":8080", nil)
}
