package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type ApiResult struct {
	Centers []*struct {
		Name     string
		Sessions []*Session
	}
}

type Session struct {
	Date          string
	Vaccine       string
	MinimumAge    int  `json:"min_age_limit"`
	MaximumAge    int  `json:"max_age_limit"`
	AllowAllAge   bool `json:"allow_all_age"`
	CapacityDose1 int  `json:"available_capacity_dose1"`
	CapacityDose2 int  `json:"available_capacity_dose2"`
}

type ValidSession struct {
	Hospital   string
	Date       string
	Vaccine    string
	MinimumAge int
	MaximumAge int
	AgeBracket string
	CapacityDose1 int
	CapacityDose2 int
}

func main() {
	districtId := flag.Int("district", 307, "id of district whose slot details should be shown")
	flag.Parse()
	results, err := GetResults(*districtId, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	for _, session := range GetSessionsWithSlots(results) {
		fmt.Println(session)
	}
}

// GetResults calls the API and returns the result.
func GetResults(districtId int, date time.Time) (*ApiResult, error) {
	resp, err := http.Get(GenerateApiUrl(districtId, date))
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

// Get only those sessions which have available slots.
func GetSessionsWithSlots(result *ApiResult) []ValidSession {
	var validSessions []ValidSession
	for _, center := range result.Centers {
		for _, session := range center.Sessions {
			if session.CapacityDose1 > 0 || session.CapacityDose2 > 0 {
				validSessions = append(validSessions, ValidSession{
					Hospital:      center.Name,
					Date:          session.Date,
					Vaccine:       session.Vaccine,
					CapacityDose1: session.CapacityDose1,
					CapacityDose2: session.CapacityDose2,
					MinimumAge:    session.MinimumAge,
					MaximumAge:    session.MaximumAge,
					AgeBracket: GetAgeBracket(*session),
				})
			}
		}
	}
	return validSessions
}

// GenerateApiUrl generates the URL that must be used
// for the given district id and date.
func GenerateApiUrl(districtId int, date time.Time) string {
	const BASE_URL_CALENDAR_BY_DISTRICT = "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict"

	// there's probably a better way to add query params to url
	v := url.Values{}
	v.Set("district_id", strconv.Itoa(districtId))
	v.Set("date", date.Format("02-01-2006"))

	return BASE_URL_CALENDAR_BY_DISTRICT + "?" + v.Encode()
}

// Get age bracket as string, given a Session.
func GetAgeBracket(s Session) string {
	if s.AllowAllAge {
		return fmt.Sprintf("%d+", s.MinimumAge)
	} else {
		return fmt.Sprintf("%d-%d", s.MinimumAge, s.MaximumAge)
	}
}

// Method to print a ValidSession.
func (vs ValidSession) String() string {
	return fmt.Sprintf("%-40s %s %-12s %-5.5s %4d %4d", vs.Hospital, vs.Date, vs.Vaccine, vs.AgeBracket, vs.CapacityDose1, vs.CapacityDose2)
}
