package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/lizrice/secure-connections/utils"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	//"net/http"
)

type RequestBody struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

func main() {
	m := flag.Int("m", 0, "input money")
	k := flag.String("k", "", "input candy type")
	c := flag.Int("c", 0, "input count candy")
	flag.Parse()

	var request RequestBody
	request.Money = *m
	request.CandyType = *k
	request.CandyCount = *c

	marsh, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error Marshal!")
	}

	client := getClient()
	resp, err := client.Post("https://localhost:3333/buy_candy", "application/json", bytes.NewBuffer(marsh))
	//resp, err := http.Post("https://localhost:8080/buy_candy", "application/json", bytes.NewBuffer(marsh))
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
}

func getClient() *http.Client {
	//file, err := os.Open("../ca/minica.pem")
	data, err := ioutil.ReadFile("../ca/minica.pem")
	if err != nil {
		fmt.Println("Certificate not open!")
		os.Exit(1)
	}
	cp, _ := x509.SystemCertPool()
	cp.AppendCertsFromPEM(data)

	config := &tls.Config{
		RootCAs:               cp,
		GetClientCertificate:  utils.ClientCertReqFunc("", ""),
		VerifyPeerCertificate: utils.CertificateChains,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}
	return client
}
