package collyclient

import (
	"fmt"
	"parallel/assets"
	"parallel/db"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func Initcollyclient_Agency3() {
	const agencyname = "DelNorte"
	var ScrapedAsset assets.Asset
	fmt.Println("Collect start at: " + time.Now().Format(time.Stamp))
	cLinks := colly.NewCollector(
		colly.AllowedDomains("www.arrendamientosdelnorte.com"),
	)
	cDetails := colly.NewCollector(
		colly.AllowedDomains("www.arrendamientosdelnorte.com"),
	)
	cLinks.Limit(&colly.LimitRule{
		DomainGlob: "*www.arrendamientosdelnorte.com*",
	})
	cDetails.Limit(&colly.LimitRule{
		DomainGlob: "*www.arrendamientosdelnorte.com*",
	})
	cDetails.OnHTML("#top-inm > div > div.col-xs-12.col-md-4.col-md-offset-5 > span", func(e *colly.HTMLElement) {
		ScrapedAsset.Business = e.Text[:11]
	})
	cDetails.OnHTML("#detalle > div > div.col-xs-12.col-md-7", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, a *colly.HTMLElement) {
			Texts := strings.Split(strings.TrimSpace(a.Text), " ")
			if len(Texts) > 1 {
				Value := strings.TrimSpace(strings.Join(Texts[1:], ""))
				switch strings.TrimSpace(Texts[0]) {
				case "Valor:":
					ScrapedAsset.Price = Value
				case "Área:":
					ScrapedAsset.Area = Value
				case "Alcobas:":
					ScrapedAsset.Numrooms = Value
				case "Baños:":
					ScrapedAsset.Numbaths = Value
				case "Barrio:":
					ScrapedAsset.Location = Value
				//case "Dirección:":
				//	ScrapedAsset.Location += Value + ", "
				case "Tipo:":
					ScrapedAsset.Type = Value
				case "Parqueadero:":
					ScrapedAsset.Parking = Value
				case "Municipio:":
					ScrapedAsset.City = Value
				}
			}

		})
		//fmt.Println(e.ChildText("b"))

	})

	cDetails.OnScraped(func(r *colly.Response) {
		if strings.HasPrefix(r.Request.URL.String(), "http://www.arrendamientosdelnorte.com") {
			SplittedURL := strings.Split(r.Request.URL.String(), "/")
			if len(SplittedURL) >= 5 {
				ScrapedAsset.Code = SplittedURL[4]
			}
			ScrapedAsset.Status = true
			ScrapedAsset.Agency = agencyname
			ScrapedAsset.Link = r.Request.URL.String()
			fmt.Println("Inserted: " + ScrapedAsset.GetCode() + " Object: " + ScrapedAsset.ToJSON())
			//db.DBInsertPostgres(RowCode,"Santafe", RowSector, RowPrice, RowArea, RowNumrooms, RowNumbaths, r.Request.URL.String(), "Active")
			db.DBInsertRedis(ScrapedAsset.GetCode(), ScrapedAsset.ToJSON())
		}
	})

	// On every a element which has href attribute call callback
	cLinks.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		//fmt.Println(link)
		if strings.HasPrefix(link, "http://www.arrendamientosdelnorte.com/") && link != "http://www.arrendamientosdelnorte.com/" && !strings.HasPrefix(link, "http://www.arrendamientosdelnorte.com/buscador.php") {
			cDetails.Visit(link)
		} else {
			return
		}
	})

	cLinks.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	cDetails.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	cLinks.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		fmt.Println("RETRYING")
		cLinks.Visit(r.Request.URL.String())
	})
	cDetails.OnResponse(func(r *colly.Response) {
		//fmt.Println(string(r.Body))
	})
	cLinks.OnResponse(func(r *colly.Response) {
		//fmt.Println(string(r.Body))
	})
	cLinks.Visit("http://www.arrendamientosdelnorte.com/buscador.php?concepto=Arriendo")

	fmt.Println("Collect end at: " + time.Now().Format(time.Stamp))
}

/*
func main() {
	Initcollyclient_Agency3()
}
*/
