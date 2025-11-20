package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	router := chi.NewRouter()
	router.Get("/", root)

	log.Printf("Starting server on port %s", os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)
}

type Animal struct {
	Color string
	Paws  int
}

func root(w http.ResponseWriter, r *http.Request) {
	animal := Animal{
		Color: "brown",
		Paws:  4,
	}
	a, _ := json.Marshal(animal)
	w.Write(a)
}
