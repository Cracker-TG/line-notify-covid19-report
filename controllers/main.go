package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/Cracker-TG/line-notify-covid19-report/forms"
	"github.com/Cracker-TG/line-notify-covid19-report/helpers"
	"github.com/enescakir/emoji"
	"github.com/gin-gonic/gin"
)

type MainController struct{}

var dataForm = new(forms.BodyData)
var h helpers.HelpersInteface = new(helpers.Helpers);

type ResponseLine struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

func (mc MainController) PushNoti(c *gin.Context) {
	key := os.Getenv("COVID_KEY")
	token_line := os.Getenv("COVID_TOKEN_LINE_NOTI")

	err := h.GetJson("https://api-lab-covid.mindbase.co.th/v2/stats/daily?key=" + key, dataForm)
	if err != nil {
		c.JSON(500, gin.H{"status": err})
	}

	endpoint := "https://notify-api.line.me/api/notify"
	data := url.Values{}
	
	data.Set("message", "โควิดวันนี้ " + h.FormatCommas(dataForm.Data.DailyCovidCases) + " " + string(emoji.PileOfPoo))

	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))

	if err != nil {
			log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	r.Header.Add("Authorization", "Bearer " + token_line)

	res, err := client.Do(r)
	if err != nil {
			log.Fatal(err)
	}

	defer res.Body.Close()

	result := new(ResponseLine)
	json.NewDecoder(res.Body).Decode(result)

	c.JSON(200, gin.H{"status": result.Status, "message": result.Message})
}
