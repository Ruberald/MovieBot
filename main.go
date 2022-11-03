package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var key, _ = os.ReadFile(".apikey")
var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("Enter the name of a film or show:")

	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(strings.Replace(name, " ", "+", -1))
	fmt.Println(name)

	resp, _ := http.Get("http://www.omdbapi.com/?apikey=" + string(key) + "&t=" + name)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	  log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)
}