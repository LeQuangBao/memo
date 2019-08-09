package main

import (
	"encoding/json"
	"net/http"
	"flag"
	"fmt"
	"time"
	"log"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Memo struct {
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
}

func main() {
	port := *flag.String("port", "7777", "Serve on port")
	// mongo
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Panic("Cannot connected to the MongoDB.")
	}
	// -- mongo
	fmt.Println("Serve on port: ", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

func setupRouting() {
	http.HandleFunc("/memo", showMemo)
	http.HandleFunc("/memos", showMemos)
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