package collyclient

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)


func Initcollyclient() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.arrendamientossantafe.com"),

	)

	c.OnHTML("li", func(e *colly.HTMLElement) {
		//		li := e.Attr("li")
		TextTitle := strings.TrimSpace(e.ChildText("b.col_50"))
		if TextTitle == "Código" || TextTitle == "Sector" || TextTitle == "Área" || TextTitle == "Precio" {
			TextValue := strings.TrimSpace(e.ChildText("div.col_50"))
			fmt.Println("Record found: " + TextTitle + ":"  + TextValue)
			fmt.Println("Actual page is: " + e.Request.URL.String())
		}
		
		//fmt.Println("Li Li: " + li)
		/*if strings.Contains(e.Text,"Sector") {
			fmt.Println("Found: " + e.Text + ":" + li)
		}
		if strings.Contains(e.Text,"Precio") {
			fmt.Println("Found: " + e.Text + ":" + li)
		}
		if strings.Contains(e.Text,"Área") {
			fmt.Println("Found: " + e.Text + ":" + li)
		}
		if strings.Contains(e.Text,"Código") {
			fmt.Println("Found: " + e.Text + ":" + li)
		}*/
	})


	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		if strings.HasPrefix(link, "/webs/santafe/pages/basico") || strings.HasPrefix(link, "/webs/santafe/inmueble")  {
			c.Visit(e.Request.AbsoluteURL(link))
		}else{
			return
		}

		
	})



	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("http://www.arrendamientossantafe.com/webs/santafe/pages/basico?bussines_type=Arrendar")
}

