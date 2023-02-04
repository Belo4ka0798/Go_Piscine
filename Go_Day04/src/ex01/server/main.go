package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "cow.h"
import "C"
import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/lizrice/secure-connections/utils"
	"net/http"
	"os"
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
	server := getServer()
	http.HandleFunc("/buy_candy", BuyCandy)
	err := server.ListenAndServeTLS("", "")
	//err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
		os.Exit(1)
	}
}

func getServer() *http.Server {

	tls := &tls.Config{
		GetCertificate:        utils.CertReqFunc("cert.pem", "key.pem"),
		VerifyPeerCertificate: utils.CertificateChains,
	}

	server := &http.Server{
		Addr:      ":3333",
		TLSConfig: tls,
	}
	return server
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
		cStr := C.CString("Hello from stdio")
		//defer C.free(unsafe.Pointer(cStr))
		b := (C.GoString)(C.ask_cow(cStr))
		json.NewEncoder(w).Encode(SuccessBuy{
			//Thanks: "Thank you!",
			Thanks: b,
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
