package main

import (
	"net/http"
	"os"
	"log"
	"github.com/gorilla/mux"
)
func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == ""{
	  port = "8000"
	}
	return ":" + port, nil
  }

  func main(){
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")
	http.Handle("/",r)
	if err := http.ListenAndServe(addr, nil);
	err != nil {
		panic(err)
	  }
  }

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}