package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var token string

func main() {
	token = os.Getenv("lichesstoken")
	getgames("genghisjahn")
}

func getgames(username string) {
	client := http.Client{Timeout: time.Second * 5}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://lichess.org/api/games/user/%v?max=10&perftype=blitz&analysed=true&clocks=true&evals=true&opening=true", username), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, _ := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
