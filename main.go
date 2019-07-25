package main

import (
	"net/http"
	"flag"
	"fmt"
)

func main() {
	port := *flag.String("port", "7777", "Serve on port")

	http.HandleFunc("/memo", showMemo)
	fmt.Println("Serve on port: ", port)
	http.ListenAndServe(":" + port, nil)
}

func showMemo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Memo"))
}