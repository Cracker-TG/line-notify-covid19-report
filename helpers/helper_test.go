package helpers_test

import (
	"testing"

	"github.com/Cracker-TG/line-notify-covid19-report/helpers"
)

var h helpers.HelpersInteface = new(helpers.Helpers)
func TestGetFormatDateThai(t * testing.T){
	date := "2021-07-22"
	f :=	h.GetFormatDateThai(date)
	if(f != "22 กรกฎาคม 2021"){
		t.Error(f)
	}
}