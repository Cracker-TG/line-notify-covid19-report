package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"
)

type HelpersInteface interface {
	GetJson(url string, target interface{})  error
    FormatCommas(num int) string
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


func (h Helpers) FormatCommas(num int) string {
    str := fmt.Sprintf("%d", num)
    re := regexp.MustCompile("(\\d+)(\\d{3})")
    for n := ""; n != str; {
        n = str
        str = re.ReplaceAllString(str, "$1,$2")
    }

    return str
}
