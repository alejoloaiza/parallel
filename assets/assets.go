package assets

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"encoding/json"
	"encoding/hex"
)
type Asset struct {
	Business string
	Code string
	Type string
	Agency string
	Sector string
	Area string
	Price string
	Numrooms string
	Numbaths string
	Status bool
	Link string
}

func  (a *Asset) GetCode() string {
	h := sha256.New()
	h.Write([]byte(a.ToString()))
	//fmt.Printf("%x", h.Sum(nil))
	return hex.EncodeToString(h.Sum(nil))
}

func  (a *Asset) ToJSON() string {
	b, err := json.Marshal(a)
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}

func (a *Asset) ToString() string{
	var AssetString string
	AssetString =  a.Business
	AssetString = AssetString + "|" + a.Code
	AssetString = AssetString + "|" + a.Type
	AssetString = AssetString + "|" + a.Agency
	AssetString = AssetString + "|" + a.Sector
	AssetString = AssetString + "|" + a.Area
	AssetString = AssetString + "|" + a.Price
	AssetString = AssetString + "|" + a.Numrooms
	AssetString = AssetString + "|" + a.Numbaths
	AssetString = AssetString + "|" + strconv.FormatBool(a.Status)
	AssetString = AssetString + "|" + a.Link
	return AssetString
}
