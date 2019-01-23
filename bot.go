package main

import (
	"./payload"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//	"strconv"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Requested")
	if req.Method != "POST" {
		fmt.Fprintf(w, "Hello, World\n")
		return
	}

	if req.Header.Get("Content-Type") != "application/json" {
		fmt.Println("Content-Type is not application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		fmt.Printf("body error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var jsonBody payload.PullRequestPayload
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		fmt.Printf("jsonBody error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("%v\n", jsonBody)

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
