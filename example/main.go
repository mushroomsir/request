package main

import (
	"github.com/mushroomsir/request"
)

func main() {
	url := "https://github.com"

	body := struct {
		Any interface{} `json:"Any"`
	}{
		Any: "",
	}

	result := struct {
		Any interface{} `json:"Any"`
	}{}

	request.Post(url).Body(&body).Result(&result).Do()
}
