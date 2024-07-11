package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type People struct {
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	passportSerieStr := r.URL.Query().Get("passportSerie")
	passportNumberStr := r.URL.Query().Get("passportNumber")

	if passportSerieStr == "" || passportNumberStr == "" {
		http.Error(w, "Missing query parameters", http.StatusBadRequest)
		return
	}

	response := People{
		Surname:    "Иванов",
		Name:       "Иван",
		Patronymic: "Иванович",
		Address:    "г. Москва, ул. Ленина, д. 5, кв. 1",
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/info", infoHandler)
	fmt.Println("Starting server at :8888")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
