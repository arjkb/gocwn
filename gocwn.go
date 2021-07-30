package main

import (
	"fmt"
	// "net/http"
	"net/url"
	"strconv"
	"time"
)

const BASE_URL = "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict"

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

func main() {
	fmt.Println("hello")
}
