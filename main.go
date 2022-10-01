package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Url struct {
	Url string `json:"url"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to supply a URL to shorten!")	
	}

	shortenUrl(os.Args[1])
}

func shortenUrl(url string) {
	client := &http.Client{}

	data := []byte(
		fmt.Sprintf(`{
			"url": "%s"
		}`, url),
	)

	req, err := http.NewRequest(
		"POST",
		"https://url.fard.rocks/shorten",
		bytes.NewBuffer(data),
	)

	if err != nil {
		log.Fatal(fmt.Sprintf("There was an error shortening the URL: %s", err))
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(fmt.Sprintf("There was an error shortening the URL: %s", err))
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(fmt.Sprintf("There was an error shortening the URL: %s", err))
	}

	var shortUrl Url

	if jsonErr := json.Unmarshal(body, &shortUrl); err != nil {
		log.Fatal(jsonErr)
	}

	res.Body.Close()

	fmt.Printf("Your shortened URL is: %s\n", shortUrl.Url)
}
