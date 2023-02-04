package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type SuccessBuy struct { // 201 status
	Thanks string
	Change int
}

type FailBuy struct { // 400 status
	Error string
}

type NoMoney struct { // 401 status
	Error string
}

type RequestBody struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

var CandyTypes = map[string]int{
	"CK": 10,
	"AA": 15,
	"NT": 17,
	"DE": 21,
	"YR": 23,
}

func main() {
	http.HandleFunc("/buy_candy", BuyCandy)

	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func BuyCandy(w http.ResponseWriter, r *http.Request) {
	var Request RequestBody
	err := json.NewDecoder(r.Body).Decode(&Request)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(FailBuy{
			Error: "Wrong request",
		})
		return
	}
	switch {
	case CandyTypes[Request.CandyType] == 0:
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(FailBuy{Error: fmt.Sprintf("No such candies!")})
		return
	case Request.Money < 0:
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(FailBuy{Error: fmt.Sprintf("Need some money!")})
		return
	}

	switch {
	case CandyTypes[Request.CandyType]*Request.CandyCount <= Request.Money:
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(SuccessBuy{
			Thanks: "Thank you!",
			Change: Request.Money - CandyTypes[Request.CandyType]*Request.CandyCount,
		})
		return
	case CandyTypes[Request.CandyType]*Request.CandyCount > Request.Money:
		w.WriteHeader(402)
		json.NewEncoder(w).Encode(NoMoney{Error: fmt.Sprintf("You need %d more money",
			CandyTypes[Request.CandyType]*Request.CandyCount-Request.Money)})
		return
	}
}
