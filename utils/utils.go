package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"parallel/config"
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

func NormalizeLocation(InputString string) (float64, float64) {

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