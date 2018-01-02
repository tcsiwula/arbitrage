// run: go run server.go

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Printf("Go server \n")
	fmt.Printf("endpoint exposed: http://localhost:8080/go \n")
	http.HandleFunc("/go", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Printf("request received at: %d\n", t)
	fmt.Fprintf(w, "go server running on port localhost:8080\n")
	fmt.Fprintf(w, "request received at: %d\n", t)

}
