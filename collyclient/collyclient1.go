package collyclient

import (
	"fmt"
	"parallel/assets"
	"parallel/db"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func Initcollyclient_Agency1() {
	const agencyname = "Santafe"
	var ScrapedAsset assets.Asset
	fmt.Println("Collect start at: " + time.Now().Format(time.Stamp))
	cLinks := colly.NewCollector(
		colly.AllowedDomains("www.arrendamientossantafe.com"),
	)
	cDetails := colly.NewCollector(
		colly.AllowedDomains("www.arrendamientossantafe.com"),
		colly.Async(true),
	)
	cLinks.Limit(&colly.LimitRule{
		DomainGlob:  "*www.arrendamientossantafe.com*",
		Delay:       4 * time.Second,
		RandomDelay: 2 * time.Second,
	})
	cDetails.Limit(&colly.LimitRule{
		DomainGlob:  "*www.arrendamientossantafe.com*",
		Parallelism: 5,
		Delay:       1 * time.Second,
		RandomDelay: 1 * time.Second,
	})
	cDetails.OnHTML("li", func(e *colly.HTMLElement) {
		TextTitle := strings.TrimSpace(e.ChildText("b.col_50"))
		switch TextTitle {
		case "Código":
			ScrapedAsset.Code = e.ChildText("div.col_50")
		case "Sector":
			ScrapedAsset.Sector = e.ChildText("div.col_50")
		case "Área":
			ScrapedAsset.Area = e.ChildText("div.col_50")
		case "Precio":
			ScrapedAsset.Price = e.ChildText("div.col_50")
		case "Nº de alcobas":
			ScrapedAsset.Numrooms = e.ChildText("div.col_50")
		case "Nº de baños":
			ScrapedAsset.Numbaths = e.ChildText("div.col_50")
		case "Tipo de inmueble":
			ScrapedAsset.Business = e.ChildText("div.col_50")
		}
		TextTitle = strings.TrimSpace(e.ChildText("a"))
		if strings.Contains(TextTitle, "- Código") {
			ScrapedAsset.Type = TextTitle
		}
	})

	cDetails.OnScraped(func(r *colly.Response) {
		if strings.HasPrefix(r.Request.URL.String(), "http://www.arrendamientossantafe.com/webs/santafe/inmueble") {
			//fmt.Println("Finished ", r.Request.URL)
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
		if strings.HasPrefix(link, "/webs/santafe/inmueble") {
			cDetails.Visit(e.Request.AbsoluteURL(link))
		} else if strings.HasPrefix(link, "/webs/santafe/pages/basico") {
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
	cLinks.Visit("http://www.arrendamientossantafe.com/webs/santafe/pages/basico?bussines_type=Arrendar")

	fmt.Println("Collect end at: " + time.Now().Format(time.Stamp))
}
