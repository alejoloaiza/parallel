package collyclient

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	//"parallel/db"
	"github.com/gocolly/redisstorage"
)


func Initcollyclient_Agency1() {

	var RowSector string
	var RowCode string
	var RowArea string
	var RowPrice string
	var RowNumrooms string
	var RowNumbaths string

	c := colly.NewCollector(
		colly.AllowedDomains("www.arrendamientossantafe.com"),

	)

	storage := &redisstorage.Storage{
		Address:  "127.0.0.1:6379",
		Password: "",
		DB:       0,
		Prefix:   "job01",
	}

	err := c.SetStorage(storage)
	if err != nil {
		panic(err)
	}
	c.OnHTML("li", func(e *colly.HTMLElement) {
		TextTitle := strings.TrimSpace(e.ChildText("b.col_50"))
		switch TextTitle {
		case "Código": RowCode=e.ChildText("div.col_50")
		case "Sector": RowSector=e.ChildText("div.col_50")
		case "Área": RowArea=e.ChildText("div.col_50")
		case "Precio": RowPrice=e.ChildText("div.col_50")
		case "Nº de alcobas": RowNumrooms=e.ChildText("div.col_50")
		case "Nº de baños": RowNumbaths=e.ChildText("div.col_50")
		}
		
	})

	c.OnScraped(func(r *colly.Response) {
		//fmt.Println(r.Request.URL.String())
		if strings.HasPrefix(r.Request.URL.String(), "http://www.arrendamientossantafe.com/webs/santafe/inmueble") {
		//fmt.Println("Finished", r.Request.URL)
			fmt.Printf("Code %s Sector %s Area %s Price %s Rooms %s Baths %s ",RowCode,RowSector,RowArea,RowPrice,RowNumrooms,RowNumbaths)
			fmt.Println("Finished ", r.Request.URL)
			//db.DBInsertPostgres(RowCode,"Santafe", RowSector, RowPrice, RowArea, RowNumrooms, RowNumbaths, r.Request.URL.String(), "Active")
		}
	})

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasPrefix(link, "/webs/santafe/pages/basico") || strings.HasPrefix(link, "/webs/santafe/inmueble")  {
			fmt.Println("LINK>> "+e.Request.AbsoluteURL(link)) 
			c.Visit(e.Request.AbsoluteURL(link))
		}else{
			return
		}
	})
/*
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
*/
	c.Visit("http://www.arrendamientossantafe.com/webs/santafe/pages/basico?bussines_type=Arrendar")
}

