package main

import (
	"testing"
	"time"
)

//
func TestGenerateApiUrl10Nov2009(t *testing.T) {
	generatedUrl := GenerateApiUrl(500, time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC))
	expectedUrl := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict?date=10-11-2009&district_id=500"
	if generatedUrl != expectedUrl {
		t.Fatalf(`Expected %s, got %s`, expectedUrl, generatedUrl)
	}
}

func TestGenerateApiUrl1Mar2020(t *testing.T) {
	generatedUrl := GenerateApiUrl(500, time.Date(2020, time.March, 1, 0, 0, 0, 0, time.UTC))
	expectedUrl := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict?date=01-03-2020&district_id=500"
	if generatedUrl != expectedUrl {
		t.Fatalf(`Expected %s, got %s`, expectedUrl, generatedUrl)
	}
}

func TestGenerateApiUrl01Mar2020(t *testing.T) {
	generatedUrl := GenerateApiUrl(500, time.Date(2020, time.March, 01, 0, 0, 0, 0, time.UTC))
	expectedUrl := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict?date=01-03-2020&district_id=500"
	if generatedUrl != expectedUrl {
		t.Fatalf(`Expected %s, got %s`, expectedUrl, generatedUrl)
	}
}
