package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	details "github.com/GorgiBytes/go-microservices/details"
	"github.com/gorilla/mux"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("In root now")
	response := map[string]int{
		"Status": http.StatusFound,
	}
	json.NewEncoder(w).Encode(response)
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Displaying Time")
	response := map[string]string{
		"Hello": "Sunshine",
		"Time":  time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func IPHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("In IP now")
	hostname, err := details.GetHostname()
	if err != nil {
		panic(err)
	}
	IP, _ := details.GetIP()
	fmt.Println(hostname, IP)
	response := map[string]string{
		"HostName": hostname,
		"IP":       IP.String(),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/time", TimeHandler)
	r.HandleFunc("/IP", IPHandler)
	http.ListenAndServe(":8080", r)
}
