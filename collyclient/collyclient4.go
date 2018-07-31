//package collyclient
package main

import (
	"fmt"
	"os"
	"parallel/assets"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func Initcollyclient_Agency4() {
	const agencyname = "EspacioUrbano"
	var ScrapedAsset assets.Asset
	fmt.Println("Collect start at: " + time.Now().Format(time.Stamp))
	cLinks := colly.NewCollector(
		colly.AllowedDomains("www.espaciourbano.com"),
	)
	cDetails := colly.NewCollector(
		colly.AllowedDomains("www.espaciourbano.com"),
		colly.Async(false)
	)
	cLinks.Limit(&colly.LimitRule{
		DomainGlob: "*www.espaciourbano.com*",
		Delay:      2 * time.Second,
		//RandomDelay: 2 * time.Second,
	})
	cLinks.SetRequestTimeout(time.Second * 60)
	cDetails.SetRequestTimeout(time.Second * 60)

	cDetails.Limit(&colly.LimitRule{
		DomainGlob: "*www.espaciourbano.com*",
		//Parallelism: 1,
		Delay: 1 * time.Second,
		//RandomDelay: 1 * time.Second,
	})
	cDetails.OnHTML("#ribbon > table > tbody > tr:nth-child(2) > td > strong.Tituloanuncio", func(e *colly.HTMLElement) {
		ScrapedAsset.Price = e.Text
		//fmt.Println("Precio>>>> " + ScrapedAsset.Price)
	})
	cDetails.OnHTML("#ribbon > table > tbody > tr:nth-child(9) > td > table > tbody > tr:nth-child(3) > td > table > tbody > tr", func(e *colly.HTMLElement) {
		if strings.TrimSpace(e.ChildText("td:nth-child(1) > span")) == "" {
			return
		}
		if strings.TrimSpace(e.ChildText("td:nth-child(2) > h1 > span")) != "" {
			ScrapedAsset.Location = e.ChildText("td:nth-child(2) > h1 > span")
		}
		//fmt.Printf("%s => %s \n", e.ChildText("td:nth-child(1) > span"), e.ChildText("td:nth-child(2) > span"))

		TextTitle := strings.TrimSpace(e.ChildText("td:nth-child(1) > span"))
		regex, _ := regexp.Compile("\n")
		TextValue := regex.ReplaceAllString(e.ChildText("td:nth-child(2) > span"), "")
		TextValue = strings.TrimSpace(TextValue)
		switch TextTitle {
		case "Alcobas familiares :":
			ScrapedAsset.Numrooms = strings.Replace(TextValue, " ", "", -1)
		case "Baños familiares :":
			ScrapedAsset.Numbaths = strings.Replace(TextValue, " ", "", -1)
		case "Area total :":
			ScrapedAsset.Area = TextValue
		case "Parqueaderos :":
			ScrapedAsset.Parking = strings.Replace(TextValue, " ", "", -1)
		}

	})
	cDetails.OnRequest(func(r *colly.Request) {
		ScrapedAsset = assets.Asset{}
	})
	cDetails.OnScraped(func(r *colly.Response) {
		if strings.HasPrefix(r.Request.URL.String(), "https://www.espaciourbano.com/BR_fichaDetalle_Vivienda") {
			//fmt.Println("Finished ", r.Request.URL)
			ScrapedAsset.Status = true
			ScrapedAsset.Agency = agencyname
			ScrapedAsset.Link = r.Request.URL.String()
			newline := fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s;%s", ScrapedAsset.Price, ScrapedAsset.Agency, ScrapedAsset.Location, ScrapedAsset.Numrooms, ScrapedAsset.Numbaths, ScrapedAsset.Area, ScrapedAsset.Parking, ScrapedAsset.Link)
			fmt.Println(newline)
			//ScrapedAsset = assets.Asset{}
			WriteResults("results.csv", newline)
			//fmt.Println("Inserted: " + ScrapedAsset.GetCode() + " Object: " + ScrapedAsset.ToJSON())
			//db.DBInsertPostgres(RowCode,"Santafe", RowSector, RowPrice, RowArea, RowNumrooms, RowNumbaths, r.Request.URL.String(), "Active")
			//db.DBInsertRedis(ScrapedAsset.GetCode(), ScrapedAsset.ToJSON())
		}
	})

	// On every a element which has href attribute call callback
	cLinks.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasPrefix(link, "BR_fichaDetalle_Vivienda") {
			//fmt.Println(e.Request.AbsoluteURL(link))

			cDetails.Visit(e.Request.AbsoluteURL(link))
		} else if strings.HasPrefix(link, "BR_Buscaranuncios_Resultado_Arriendo") {
			//fmt.Println(e.Request.AbsoluteURL(link))

			cLinks.Visit(e.Request.AbsoluteURL(link))
		} else {
			return
		}
	})

	cLinks.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	cLinks.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		fmt.Println("RETRYING")
		cLinks.Visit(r.Request.URL.String())
	})
	cLinks.Visit("https://www.espaciourbano.com/BR_BuscarAnuncios_Resultado_Arriendo.asp?sNegocio=1&sZona=1&sCiudad=10007&sRango1=0&sRango2=999999999999&button=Buscar")

	fmt.Println("Collect end at: " + time.Now().Format(time.Stamp))
}
func main() {
	Initcollyclient_Agency4()
}

func WriteResults(logpath string, line string) {
	f, err := os.OpenFile(logpath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	newline := line + string('\n')
	_, err = f.WriteString(newline)
	if err != nil {
		panic(err)
	}
}
