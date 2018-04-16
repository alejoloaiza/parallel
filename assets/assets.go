package assets

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Asset struct {
	Business string  `json:"business"`
	Code     string  `json:"code"`
	Type     string  `json:"type"`
	Agency   string  `json:"agency"`
	Location string  `json:"location"`
	City     string  `json:"city"`
	Area     string  `json:"area"`
	Price    string  `json:"price"`
	Numrooms string  `json:"numrooms"`
	Numbaths string  `json:"numbaths"`
	Status   bool    `json:"status"`
	Link     string  `json:"link"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
}

func (a *Asset) GetCode() string {
	h := sha256.New()
	h.Write([]byte(a.ToJSON()))
	//fmt.Printf("%x", h.Sum(nil))
	return hex.EncodeToString(h.Sum(nil))
}

func (a *Asset) ToJSON() string {
	b, err := json.Marshal(a)
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}
func (a *Asset) FromJSON(s string) {
	input := []byte(s)
	err := json.Unmarshal(input, a)
	if err != nil {
		fmt.Println("error:", err)
	}
}

/*
func (a *Asset) ToString() string {
	var AssetString string
	AssetString = a.Business
	AssetString = AssetString + "|" + a.Code
	AssetString = AssetString + "|" + a.Type
	AssetString = AssetString + "|" + a.Agency
	AssetString = AssetString + "|" + a.Location
	AssetString = AssetString + "|" + a.City
	AssetString = AssetString + "|" + a.Area
	AssetString = AssetString + "|" + a.Price
	AssetString = AssetString + "|" + a.Numrooms
	AssetString = AssetString + "|" + a.Numbaths
	AssetString = AssetString + "|" + strconv.FormatBool(a.Status)
	AssetString = AssetString + "|" + a.Link
	AssetString = AssetString + "|" + strconv.FormatFloat(a.Lat, 'E', -1, 64)
	AssetString = AssetString + "|" + strconv.FormatFloat(a.Lon, 'E', -1, 64)
	return AssetString
}
*/
