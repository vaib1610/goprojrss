package main

import (
	"encoding/json"
	"log"
	"net/http"
)



func respondWithJSON(w http.ResponseWriter,code int,payload interface{}){
	dat,err := json.Marshal(payload)
	if err!=nil{
	log.Println("Failed to marshal JSON response: %v",payload)
	w.WriteHeader(500)
	return
}



w.Header().Add("Content-Type","application/json")
w.WriteHeader(code)
w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}
