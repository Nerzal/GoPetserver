package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var srv *http.Server
var currentDogs = getDogs()

type dog struct {
	Name string
	Age  int
}

func getDogs() []dog {
	var myDog = dog{"Wuffi", 3}
	var otherDog = dog{"WuffWuff", 4}
	var otherOtherDog = dog{"WuffiWuffWuff", 5}
	myDogs := []dog{myDog, otherDog, otherOtherDog}
	return myDogs
}

func postDog(w http.ResponseWriter, r *http.Request) {
	var newDog dog
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	var error = json.NewDecoder(r.Body).Decode(&newDog)
	if error != nil {
		http.Error(w, error.Error(), 400)
		return
	}

	w.Write([]byte("Der Hund wurde hinzugefügt"))
}

func getAllDogs(w http.ResponseWriter, r *http.Request) {

	var response, error = json.Marshal(currentDogs)
	if error != nil {
		panic(error)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func initializeServer() {
	r := mux.NewRouter()
	r.HandleFunc("/Api/Dogs", getAllDogs).Methods("GET")
	r.HandleFunc("/Api/Dogs", postDog).Methods("POST")
	srv = &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:443",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func main() {
	currentDogs = getDogs()
	initializeServer()
	var error = srv.ListenAndServe()
	if error != nil {
		log.Fatal(error)
	}
}
