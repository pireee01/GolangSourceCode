package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Mahasiswa struct {
	Nama   string `json:"Nama"`
	Nim    string `json:"NIM"`
	Alamat string `json:"Alamat"`
}

func main() {
	http.HandleFunc("/get", getMahasiswaHandler)
	http.HandleFunc("/post", postMahasiswaHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getMahasiswaHandler(w http.ResponseWriter, r *http.Request) {
	nama := r.URL.Query().Get("Nama")
	if nama == "" {
		http.Error(w, "Missing nama parameter", http.StatusBadRequest)
		return
	}

	mhs := Mahasiswa{
		Nama:   "Pirelli Rahelya Piri",
		Nim:    "2540131310",
		Alamat: "University Village",
	}

	jsonBytes, err := json.Marshal(mhs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func postMahasiswaHandler(w http.ResponseWriter, r *http.Request) {
	var mhs Mahasiswa

	err := json.NewDecoder(r.Body).Decode(&mhs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Received mahasiswa: %+v", mhs)

	w.WriteHeader(http.StatusCreated)
}
