package main

import (
	"encoding/json"
	"net/http"
	"flag"
	"fmt"
	"time"
)

type Memo struct {
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
}

func main() {
	port := *flag.String("port", "7777", "Serve on port")

	http.HandleFunc("/memo", showMemo)
	http.HandleFunc("/memos", showMemos)
	fmt.Println("Serve on port: ", port)
	http.ListenAndServe(":" + port, nil)
}

func showMemo(w http.ResponseWriter, r *http.Request) {
	peter := Memo{
		Title    : "John",
		Body     : "Doe",
		CreatedAt: time.Now(),
	}

	json.NewEncoder(w).Encode(peter)
}

func showMemos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("More memo"))
}