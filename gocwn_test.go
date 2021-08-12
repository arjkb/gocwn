package main

import (
	"testing"
	"time"
)

// TestGenerateApiUrl10Nov2009 creates a URL for date of 10 Nov 2009,
// checking that the URL generated for that date is correct.
func TestGenerateApiUrl10Nov2009(t *testing.T) {
	generatedUrl := GenerateApiUrl(500, time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC))
	expectedUrl := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict?date=10-11-2009&district_id=500"
	if generatedUrl != expectedUrl {
		t.Fatalf(`Expected %s, got %s`, expectedUrl, generatedUrl)
	}
}

// TestGenerateApiUrl10Nov2009 creates a URL for date of 1 Mar 2020,
// checking that the URL generated for that date is correct.
func TestGenerateApiUrl1Mar2020(t *testing.T) {
	generatedUrl := GenerateApiUrl(500, time.Date(2020, time.March, 1, 0, 0, 0, 0, time.UTC))
	expectedUrl := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict?date=01-03-2020&district_id=500"
	if generatedUrl != expectedUrl {
		t.Fatalf(`Expected %s, got %s`, expectedUrl, generatedUrl)
	}
}

// TestGenerateApiUrl10Nov2009 creates a URL for date of 01 Mar 2020,
// checking that the URL generated for that date is correct.
func TestGenerateApiUrl01Mar2020(t *testing.T) {
	generatedUrl := GenerateApiUrl(500, time.Date(2020, time.March, 01, 0, 0, 0, 0, time.UTC))
	expectedUrl := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict?date=01-03-2020&district_id=500"
	if generatedUrl != expectedUrl {
		t.Fatalf(`Expected %s, got %s`, expectedUrl, generatedUrl)
	}
}

// Test age bracket with the following params:
// MinimumAge=10, MaximumAge=20, AllowAllAge=false
func TestAgeBracketMin10Max20AllowAllFalse(t *testing.T) {
	validSession := Session{
		Date:          "1-2-2003",
		Vaccine:       "COVISHIELD",
		MinimumAge:    10,
		MaximumAge:    20,
		AllowAllAge:   false,
		CapacityDose1: 100,
		CapacityDose2: 200,
	}

	expectedOutput := "10-20"
	actualOutput := GetAgeBracket(validSession)

	if actualOutput != expectedOutput {
		t.Fatalf(`Expected %q, got %q`, expectedOutput, actualOutput)
	}
}

// Test age bracket with the following params:
// MinimumAge=10, MaximumAge=20, AllowAllAge=true
func TestAgeBracketMin10Max20AllowAllTrue(t *testing.T) {
	validSession := Session{
		Date:          "1-2-2003",
		Vaccine:       "COVISHIELD",
		MinimumAge:    10,
		MaximumAge:    20,
		AllowAllAge:   true,
		CapacityDose1: 100,
		CapacityDose2: 200,
	}

	expectedOutput := "10-20"
	actualOutput := GetAgeBracket(validSession)

	if actualOutput != expectedOutput {
		t.Fatalf(`Expected %q, got %q`, expectedOutput, actualOutput)
	}
}

// Test age bracket with the following params:
// MinimumAge=10, MaximumAge=0 (zero-value), AllowAllAge=true
func TestAgeBracketMin10Max0AllowAllTrue(t *testing.T) {
	validSession := Session{
		Date:          "1-2-2003",
		Vaccine:       "COVISHIELD",
		MinimumAge:    10,
		MaximumAge:    0,
		AllowAllAge:   true,
		CapacityDose1: 100,
		CapacityDose2: 200,
	}

	expectedOutput := "10+"
	actualOutput := GetAgeBracket(validSession)

	if actualOutput != expectedOutput {
		t.Fatalf(`Expected %q, got %q`, expectedOutput, actualOutput)
	}
}

// Test age bracket with the following params:
// MinimumAge=10, MaximumAge=0 (zero-value), AllowAllAge=false
func TestAgeBracketMin10Max0AllowAllFalse(t *testing.T) {
	validSession := Session{
		Date:          "1-2-2003",
		Vaccine:       "COVISHIELD",
		MinimumAge:    10,
		MaximumAge:    0,
		AllowAllAge:   false,
		CapacityDose1: 100,
		CapacityDose2: 200,
	}

	expectedOutput := "10+"
	actualOutput := GetAgeBracket(validSession)

	if actualOutput != expectedOutput {
		t.Fatalf(`Expected %q, got %q`, expectedOutput, actualOutput)
	}
}
