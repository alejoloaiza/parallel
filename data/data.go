package data

import (
	"encoding/json"
	"fmt"
	"parallel/assets"
	"parallel/db"
	"parallel/utils"

	"github.com/schollz/closestmatch"
)

var RawAssets []assets.Asset
var TransformedAsset []assets.Asset

func FillRawAssetsArray() {
	var TempAssets assets.Asset
	rediskeys := db.DBGetAllKeysRedis()
	for _, currentkey := range rediskeys {
		_ = json.Unmarshal([]byte(currentkey), &TempAssets)
		RawAssets = append(RawAssets, TempAssets)
		//fmt.Println(TempAssets.ToJSON())
	}
}
func PrintAssetsArray() {
	for _, curAsset := range TransformedAsset {
		fmt.Println(curAsset.ToJSON())
	}
}

// We define the Asset.Type based on similarity of words using schollz/closestmatch
func AssetClassifier() {

	TypeStuff := []string{"Apartamento", "Casa", "Bodega", "Finca", "Oficina", "Local"}
	BusinessStuff := []string{"Arrendar", "Vender"}
	bagSizes := []int{2, 3, 4, 5}

	cmType := closestmatch.New(TypeStuff, bagSizes)
	cmBusiness := closestmatch.New(BusinessStuff, bagSizes)
	for _, curAsset := range RawAssets {
		curAsset.Type = cmType.Closest(curAsset.Type)
		curAsset.Business = cmBusiness.Closest(curAsset.Business)
		curAsset.Area = utils.NormalizeArea(curAsset.Area)
		curAsset.Price = utils.NormalizeAmount(curAsset.Price)
		curAsset.Numrooms = utils.NormalizeAmount(curAsset.Numrooms)
		curAsset.Numbaths = utils.NormalizeAmount(curAsset.Numbaths)
		curAsset.Lat, curAsset.Lon = utils.NormalizeLocation(curAsset.Location)
		//fmt.Println("Lat: " + strconv.FormatFloat(curAsset.Lat, 'e', -1, 64) + " Lon: " + strconv.FormatFloat(curAsset.Lon, 'e', -1, 64))
		TransformedAsset = append(TransformedAsset, curAsset)
		//fmt.Println("Original: " + curAsset.Type + " Definido: " + cm.Closest(curAsset.Type))
	}

}

/*
func AssetBusinessClassifier() {

	TypeStuff := []string{"Arrendar", "Vender"}
	bagSizes := []int{2, 3, 4, 5}
	cmType := closestmatch.New(TypeStuff, bagSizes)
	cmBusiness := closestmatch.New(TypeStuff, bagSizes)
	for _, curAsset := range RawAssets {
		curAsset.Business = cm.Closest(curAsset.Business)
	}

		classifier := bayesian.NewClassifier(Rent, Sell)
		RentStuff := []string{"Arriendo", "Renta", "Arrienda"}
		SellStuff := []string{"Vende", "Venta", "Vendo"}
		classifier.Learn(RentStuff, Rent)
		classifier.Learn(SellStuff, Sell)
		for _, curAsset := range RawAssets {
			probs, _, _ := classifier.ProbScores(strings.Split(curAsset.Business, " "))
			for _, curprob := range probs {
				fmt.Println("Probabilidad " + strconv.FormatFloat(curprob, 'E', -1, 64))
			}
		}


	cm := closestmatch.New(TypeStuff, bagSizes)
	for _, curAsset := range RawAssets {
		curAsset.Business = cm.Closest(curAsset.Business)
	}

}
*/
