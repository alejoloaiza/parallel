package utils

import (
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

func NormalizeLocation(InputString string) (int, int) {
	apiurl := strings.Join(config.Localconfig.GoogleAPI, "") + InputString
	fmt.Println(apiurl)
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
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
	return 1, 1
}
