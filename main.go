package main

import (
	"log"
	"net/http"
	"os"
)

// TableAction describes the actions taken on a table
type TableAction struct {
	Action      string
	Description string
	Emote       string
}

var (
	tableFlip  = TableAction{Action: "flip", Description: "Table flipped!", Emote: "(╯°□°)╯︵ ┻━┻"}
	tableCatch = TableAction{Action: "catch", Description: "Table caught!", Emote: "┬─┬ノ( º _ ºノ)"}
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	log.SetOutput(os.Stdout)
	log.Println("Starting slack-table-emotes service...")

	http.HandleFunc("/healthz", healthHandler)
	http.HandleFunc("/table", tableHandler)

	port := getEnv("PORT", "8080")
	log.Printf("Server listening on port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
