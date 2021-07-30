package main

import (
	"fmt"
	// "net/http"
	"net/url"
	"strconv"
	"time"
)

type ApiResult struct {
	Centers []*Center
}

type Center struct {
	Name     string
	Sessions []*Session
}

type Session struct {
	Date          string
	MinimumAge    int  `json:"min_age_limit"`
	AllowAllAge   bool `json:"allow_all_age"`
	CapacityDose1 int
	CapacityDose2 int
}

func GenerateApiUrl(districtId int, date time.Time) string {
	const BASE_URL_CALENDAR_BY_DISTRICT = "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict"

	// there's probably a better way to add query params to url
	v := url.Values{}
	v.Set("district_id", strconv.Itoa(districtId))
	v.Set("date", date.Format("02-01-2006"))

	return BASE_URL_CALENDAR_BY_DISTRICT + "?" + v.Encode()
}

func main() {
	fmt.Println(GenerateApiUrl(307, time.Now()))
}
