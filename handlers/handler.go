package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"jgnovak.net/loanshark-api/models"
	"jgnovak.net/loanshark-api/services"
)

func loanHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var loanRequest models.LoanRequest
	err := json.NewDecoder(r.Body).Decode(&loanRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	loan := services.CalculateLoan(loanRequest)
	response, err := json.Marshal(loan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/loans", loanHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
