package forms

type BodyData struct {
	StatusCode string `json:"status_code"`
	StatusMsg string `json:"status_msg"`
	Timestamp string `json:"timestamp"`
	SpanId string `json:"span_id"`
	ErrMsg string `json:"error_msg"`
	Data struct{
		ID string `json:"id"`
		IsCurrentDate bool `json:"is_current_date"`
		DailyCovidGeneral int `json:"daily_covid_general"`
		DailyCovidPrison int `json:"daily_covid_prison"`
		DailyCovidCases int `json:"daily_covid_cases"`
		DailyRecovered  int `json:"daily_recovered"`
		DailyStayPatient int `json:"daily_stay_patient"`
		DailyDeaths int `json:"daily_deaths"`
		DailyVaccine1 int `json:"daily_vaccine1"`
		CumulativeVaccine1 int `json:"cumulative_vaccine1"`
		DailyVaccine2 int `json:"daily_vaccine2"`
		CumulativeVaccine2 int `json:"cumulative_vaccine2"`
		DailyVaccine3 int `json:"daily_vaccine3"`
		CumulativeVaccine3 int `json:"cumulative_vaccine3"`
		CumulativeCovidCases int `json:"cumulative_covid_cases"`
		CumulativeCovidCases63 int `json:"cumulative_covid_cases63"`
		CumulativeRecoveredCases int `json:"cumulative_recovered_cases"`
		CumulativeRecoveredCases63 int `json:"cumulative_recovered_cases63"`
		CumulativeVaccines int `json:"cumulative_vaccines"`
		CumulativeDeaths int `json:"cumulative_deaths"`
		CumulativeDeaths63 int `json:"cumulative_deaths63"`
		YesterDayCovidGeneral int `json:"yesterday_covid_general"`
		YesterDayCovidPrison int `json:"yesterday_covid_prison"`
		YesterDayCovidDeaths int `json:"yesterday_covid_deaths"`
		YesterDayCumulativeCovidCases int `json:"yesterday_cumulative_covid_cases"`
		YesterDayCumulativeRecoveredCases int `json:"yesterday_cumulative_recovered_cases"`
		CreatedDate string `json:"created_date"`
		CreatedBy string `json:"created_by"`
		UpdatedDate string `json:"updated_date"`
		UpdatedBy string `json:"updated_by"`
	}
}