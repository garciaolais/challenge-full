package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	c "github.com/garciaolais/challenge-backend/challengelib"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Serve(port string) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/dividend/{id}", returnDividend)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func returnDividend(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	data := []c.Data{
		{Divisor: 15, Message: "Linianos"},
		{Divisor: 3, Message: "Linio"},
		{Divisor: 5, Message: "IT"},
	}

	i, _ := strconv.Atoi(key)
	if message, ok := c.GetMessage(i, data); ok {
		json.NewEncoder(w).Encode(message)
	} else {
		json.NewEncoder(w).Encode(key)
	}
}
