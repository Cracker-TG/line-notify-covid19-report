package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type HelpersInteface interface {
	GetJson(url string, target interface{})  error
    FormatCommas(num int) string
    GetFormatDateThai(date string) string
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

func (h Helpers) GetFormatDateThai(date string) string  {
    const (layoutISO = "2006-01-02")
    month_thai := []string{"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน", "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"}
	t, _ := time.Parse(layoutISO, date)
    f := strconv.Itoa(t.Day()) + " " + month_thai[t.Month()-1] + " " + strconv.Itoa(t.Year())
    return f
}
