package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Jokedata struct {
	Category string `json: "category"`
	Joke     string `json: "joke"`
	Id       int    `json: "id"`
	Lang     string `json: "lang"`
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	url := goDotEnvVariable("API_URL")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var data Jokedata
	json.Unmarshal(content, &data)
	// json.NewDecoder(response.Body).Decode(&data) --> An alternate way to get the response data without usint ioutil
	fmt.Println(data.Joke)
}
