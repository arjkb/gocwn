package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	Vaccine       string
	MinimumAge    int  `json:"min_age_limit"`
	AllowAllAge   bool `json:"allow_all_age"`
	CapacityDose1 int  `json:"available_capacity_dose1"`
	CapacityDose2 int  `json:"available_capacity_dose2"`
}

func GenerateApiUrl(districtId int, date time.Time) string {
	const BASE_URL_CALENDAR_BY_DISTRICT = "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict"

	// there's probably a better way to add query params to url
	v := url.Values{}
	v.Set("district_id", strconv.Itoa(districtId))
	v.Set("date", date.Format("02-01-2006"))

	return BASE_URL_CALENDAR_BY_DISTRICT + "?" + v.Encode()
}

func GetResults(district_id int, date time.Time) (*ApiResult, error) {
	resp, err := http.Get(GenerateApiUrl(district_id, date))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Query failed: %s", resp.Status)
	}

	var result ApiResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func main() {
	results, err := GetResults(307, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	for _, center := range results.Centers {
		for _, session := range center.Sessions {
			fmt.Printf("%-40s %s %-12s %d %4d %4d\n", center.Name, session.Date, session.Vaccine, session.MinimumAge, session.CapacityDose1, session.CapacityDose2)
		}
	}
}