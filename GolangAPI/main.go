package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Pessoa struct {
	ID        string `json."id,omniempty"`
	Firstname string `json."firstname,omniempty"`
	Lastname  string `json."lastname,omniempty"`
	City      string `json."city,omniempty"`
}

var pessoas []Pessoa

func MostrarPessoas(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(pessoas)
}

func MostrarPessoa(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range pessoas {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Pessoa{})
}

func CriarPessoa(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var pessoa Pessoa
	_ = json.NewDecoder(req.Body).Decode(&pessoa)
	pessoa.ID = params["id"]
	pessoas = append(pessoas, pessoa)
	json.NewEncoder(w).Encode(pessoas)
}

func ApagarPessoa(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range pessoas {
		if item.ID == params["id"] { //Aqui era onde tava o erro durante a palestra, "ID" no lugar de "id"
			pessoas = append(pessoas[:index], pessoas[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(pessoas)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/pessoas", MostrarPessoas).Methods("GET")
	router.HandleFunc("/pessoas/{id}", MostrarPessoa).Methods("GET")
	router.HandleFunc("/pessoas/{id}", CriarPessoa).Methods("POST")
	router.HandleFunc("/pessoas/{id}", ApagarPessoa).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))
}
