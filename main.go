package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var version = "dev"

func main() {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "DemoService"
	}
	port := os.Getenv("APP_PORT")
	if port == "" {
		panic("You must specify a port to listen on using the APP_PORT environment variable")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Service-Name", appName)
		w.Header().Set("Service-Version", version)
		encoder := json.NewEncoder(w)
		encoder.Encode(map[string]string{
			"message": "Hello World!",
		})
	})
	fmt.Printf("%s Version %s started listening on :%s\n", appName, version, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
