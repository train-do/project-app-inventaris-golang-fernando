package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/train-do/project-app-inventaris-golang-fernando/collection"
)

func CalculateDepreciation(initPrice int, usageDays int, rateDepreciation int) int {
	depreciationFreq := usageDays / 365
	if depreciationFreq == 0 {
		return 0
	}
	result := initPrice
	for i := 0; i < depreciationFreq; i++ {
		result -= (result * rateDepreciation / 100)
	}
	return initPrice - result
}
func CalculateTotalUsageDays(timePurchase time.Time) int {
	days := time.Since(timePurchase).Hours() / 24
	return int(days)
}
func ToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}

func SetResponse(w http.ResponseWriter, success bool, data collection.Response, statusCode int, message string) collection.Response {
	w.WriteHeader(statusCode)
	data.Success = success
	data.StatusCode = statusCode
	data.Message = message
	return data
}

func ToTimeFormat(timeStr string) time.Time {
	parsedTime, err := time.Parse("2006-01-02", timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	return parsedTime
}
