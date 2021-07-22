package helpers

import (
	"encoding/json"
	"net/http"
	"time"
)

type HelpersInteface interface {
	GetJson(url string, target interface{})  error
}

type Helpers struct{}

var myClient = &http.Client{Timeout: 100 * time.Second}

func(h Helpers) GetJson(url string, target interface{})  error {
    r, err := myClient.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}