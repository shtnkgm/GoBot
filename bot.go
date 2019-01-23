package main

import (
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
	// リクエストメソッドのチェック
	if req.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// コンテンツ形式のチェック
	if req.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// コンテンツの長さの取得
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// リクエストbodyの取得
	body := make([]byte, length)
	length, err = req.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// jsonの取得
	var jsonBody map[string]interface{}
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
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
