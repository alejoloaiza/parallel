package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"parallel/config"
	"regexp"
	"strings"
)

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}
func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
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
func NormalizeParking(InputString string) string {
	TempString := strings.ToLower(InputString)
	TempString = strings.TrimSpace(TempString)
	if TempString != "" && TempString != "0" {
		return "1"
	} else {
		return "0"
	}
}
func NormalizeLocation(InputString string, Api string) (float64, float64) {
	if Api != "on" {
		return 0, 0
	}
	apiurl := strings.Join(config.Localconfig.GoogleAPI, "") + InputString
	req, err := http.NewRequest("GET", apiurl, nil)
	if err != nil {
		fmt.Println("Error in newRequest: ", err)

	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error in Response: ", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		var m = new(struct {
			Results []struct {
				Geometry struct {
					Location struct {
						Lat float64 `json:"lat"`
						Lng float64 `json:"lng"`
					}
				}
			}
		})
		var err = json.Unmarshal(bodyBytes, &m)
		//fmt.Println(err)
		var data map[string]interface{}
		err = json.Unmarshal(bodyBytes, &data)
		if err != nil {
			fmt.Println("Error in Unmarshal: ", err)
		}
		return m.Results[0].Geometry.Location.Lat, m.Results[0].Geometry.Location.Lng

	}
	return 0, 0
}
