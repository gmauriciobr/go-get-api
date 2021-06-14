package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	readFile, err := os.Open("ceps.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		consultaCep(fileScanner.Text())
	}
	readFile.Close()
}

func consultaCep(cep string) {
	var host = "https://viacep.com.br/ws/"
	resp, err := http.Get(host + cep + "/json")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}
