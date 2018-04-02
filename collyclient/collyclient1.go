package collyclient

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"parallel/db"
	"parallel/assets"
	"time"
)	


func Initcollyclient_Agency1() {
	const agencyname = "Santafe"
	var ScrapedAsset assets.Asset 
	fmt.Println("Collect start at: " + time.Now().Format(time.Stamp) )
	c := colly.NewCollector(
		colly.AllowedDomains("www.arrendamientossantafe.com"),
		colly.Async(true),

	)
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*www.arrendamientossantafe.com*",
		Parallelism: 2,
		Delay: 2 * time.Second,
		RandomDelay: 2 * time.Second,
	})

	c.OnHTML("li", func(e *colly.HTMLElement) {
		TextTitle := strings.TrimSpace(e.ChildText("b.col_50"))
		switch TextTitle {
		case "Código": ScrapedAsset.Code=e.ChildText("div.col_50")
		case "Sector": ScrapedAsset.Sector=e.ChildText("div.col_50")
		case "Área": ScrapedAsset.Area=e.ChildText("div.col_50")
		case "Precio": ScrapedAsset.Price=e.ChildText("div.col_50")
		case "Nº de alcobas": ScrapedAsset.Numrooms=e.ChildText("div.col_50")
		case "Nº de baños": ScrapedAsset.Numbaths=e.ChildText("div.col_50")
		}
		TextTitle = strings.TrimSpace(e.ChildText("a"))
		if strings.HasPrefix(TextTitle,"Apartamento") {
			ScrapedAsset.Type = "Apartment"
		} else if  strings.HasPrefix(TextTitle,"Casa") {
			ScrapedAsset.Type = "House"
		} else if  strings.HasPrefix(TextTitle,"Local") {
			ScrapedAsset.Type = "Retail"
		} else if  strings.HasPrefix(TextTitle,"Finca") {
			ScrapedAsset.Type = "Farmhouse"
		} else if  strings.HasPrefix(TextTitle,"Oficina") {
			ScrapedAsset.Type = "Office"
		} else if  strings.HasPrefix(TextTitle,"Bodega") {
			ScrapedAsset.Type = "Warehouse"
		}
	})

	c.OnScraped(func(r *colly.Response) {
		if strings.HasPrefix(r.Request.URL.String(), "http://www.arrendamientossantafe.com/webs/santafe/inmueble") {
			//fmt.Println("Finished ", r.Request.URL)
			ScrapedAsset.Status=true
			ScrapedAsset.Agency=agencyname
			ScrapedAsset.Business = "Rent"
			ScrapedAsset.Link = r.Request.URL.String()
			//fmt.Println("Inserted: " + ScrapedAsset.GetCode() + " Object: " + ScrapedAsset.ToJSON())
			//db.DBInsertPostgres(RowCode,"Santafe", RowSector, RowPrice, RowArea, RowNumrooms, RowNumbaths, r.Request.URL.String(), "Active")
			db.DBInsertRedis(ScrapedAsset.GetCode(),ScrapedAsset.ToJSON())
		}
	})

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasPrefix(link, "/webs/santafe/pages/basico") || strings.HasPrefix(link, "/webs/santafe/inmueble")  {
			//fmt.Println("LINK>> "+e.Request.AbsoluteURL(link)) 
			//time.Sleep(500 * time.Millisecond)
			c.Visit(e.Request.AbsoluteURL(link))
		}else{
			return
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("http://www.arrendamientossantafe.com/webs/santafe/pages/basico?bussines_type=Arrendar")

	fmt.Println("Collect end at: " + time.Now().Format(time.Stamp) )
}

