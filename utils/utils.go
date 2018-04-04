package utils

import (
	"regexp"
	"strings"
)

func NormalizeArea(InputString string) string {
	TempString := strings.ToLower(InputString)
	TempString = strings.Replace(TempString, "m2", "", -1)
	TempString = strings.Replace(TempString, "mt2", "", -1)
	TempString = strings.TrimSpace(TempString)
	re := regexp.MustCompile("[0-9]+")
	return strings.Join(re.FindAllString(TempString, -1), "")
}
func NormalizeAmount(InputString string) string {
	TempString := strings.ToLower(InputString)
	TempString = strings.Replace(TempString, "$", "", -1)
	TempString = strings.Replace(TempString, ",", "", -1)
	TempString = strings.Replace(TempString, ".", "", -1)
	TempString = strings.TrimSpace(TempString)
	re := regexp.MustCompile("[0-9]+")
	return strings.Join(re.FindAllString(TempString, -1), "")
}
