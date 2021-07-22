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

	if(dataForm.StatusCode == "MDB-500"){
		c.JSON(401, gin.H{"status": false, "mssage": dataForm.ErrMsg})	
		return
	}

	endpoint := "https://notify-api.line.me/api/notify"
	data := url.Values{}

	const (
    layoutISO = "2006-01-02"
    layoutUS  = "2 January 2006"
	)

	date :=  dataForm.Timestamp
	parts_date := strings.Split(date, "T")

	date_thai := h.GetFormatDateThai(parts_date[0])

	msg :=  string(emoji.FlagForThailand) + 
					"\nสถานการณ์ COVID-19 ในประเทศไทย\nวันที่  " + date_thai +
					"\nผู้ติดเชื้อเพิ่มวันนี้ +" + h.FormatCommas(dataForm.Data.DailyCovidCases) + " " + string(emoji.FaceWithMedicalMask) +
					"\nผู้ติดเชื้อใหม่ +" + h.FormatCommas(dataForm.Data.DailyCovidGeneral) + " " + string(emoji.FaceWithThermometer) +
					"\nจากเรือนจำ/ที่ต้องขัง +" + h.FormatCommas(dataForm.Data.DailyCovidPrison) + " " + string(emoji.OfficeBuilding) +
					"\nเสียชีวิตเพิ่มวันนี้ " + h.FormatCommas(dataForm.Data.DailyDeaths) + " " + string(emoji.SkullAndCrossbones) +
					"\nผู้ป่วยสะสม " + h.FormatCommas(dataForm.Data.CumulativeCovidCases) + " " +  string(emoji.Hospital) +
					"\nหายป่วยสะสม" + h.FormatCommas(dataForm.Data.CumulativeRecoveredCases) + " " + string(emoji.SlightlySmilingFace)

	data.Set("message", msg)

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
