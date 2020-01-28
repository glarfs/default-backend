package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {

	// return 404 on / path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var xcode = r.Header.Get("X-Code")
		if len(xcode) == 0 {
			xcode = "404"
		}
		xcodeint, err := strconv.Atoi(xcode)
		if err != nil {
			log.Fatal("Can't parse integer: " + xcode)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			w.WriteHeader(http.StatusNotFound)
			return			
		}
		w.WriteHeader(xcodeint)
		
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		file, errfile := ioutil.ReadFile("./assets/" + xcode + ".html")
		if errfile != nil {
			http.Error(w, errfile.Error(), http.StatusInternalServerError)
			log.Fatal("Can't find error html page: " + xcode)
			return
		}
		w.Write(file)
	})

	// return 200 on healthz path
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"status\":\"healthy!\"}"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
