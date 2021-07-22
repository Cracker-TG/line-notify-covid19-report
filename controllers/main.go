package controllers

import (
	"os"

	"github.com/Cracker-TG/line-notify-covid19-report/forms"
	"github.com/Cracker-TG/line-notify-covid19-report/helpers"
	"github.com/gin-gonic/gin"
)

type MainController struct{}

var dataForm = new(forms.BodyData)
var h helpers.HelpersInteface = new(helpers.Helpers);
func (mc MainController) PushNoti(c *gin.Context) {
	key := os.Getenv("COVID_KEY")
	err := h.GetJson("https://api-lab-covid.mindbase.co.th/v2/stats/daily?key=" + key, dataForm)
	if err != nil {
		c.JSON(500, gin.H{"status": err})
	}

	c.JSON(200, gin.H{"status": dataForm})
}
