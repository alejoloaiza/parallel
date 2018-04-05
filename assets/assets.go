package assets

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
)

type Asset struct {
	Business string
	Code     string
	Type     string
	Agency   string
	Location string
	City     string
	Area     string
	Price    string
	Numrooms string
	Numbaths string
	Status   bool
	Link     string
	Lat      float64
	Lon      float64
}

func (a *Asset) GetCode() string {
	h := sha256.New()
	h.Write([]byte(a.ToString()))
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
