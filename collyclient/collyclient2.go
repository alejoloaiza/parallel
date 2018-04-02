package collyclient

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"parallel/db"
	"parallel/assets"
	"time"
	"strconv"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"bytes"
)	

type Responsejson struct {
	Results string `json:"results"`
	ResultsCenter string `json:"results_center"`
	Print string `json:"print"`
	Order string `json:"order"`
	Lang_id string `json:"lang_id"`
	Total_Rows string `json:"total_rows"`
}

func Initcollyclient_Agency2() {
	pagenum := 0
	const agencyname = "Alnago"
	var ScrapedAsset assets.Asset 
	fmt.Println("Collect start at: " + time.Now().Format(time.Stamp) )
	c := colly.NewCollector(
		colly.AllowedDomains("www.alnago.com"),
		//colly.Async(true),
		//colly.AllowURLRevisit(),
	)
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*www.alnago.com*",
		//Parallelism: 1,
		Delay: 4 * time.Second,
		RandomDelay: 2 * time.Second,
	})


	c.OnScraped(func(r *colly.Response) { 
		fmt.Println("Finished ", r.Request.URL) 
		if strings.HasPrefix(r.Request.URL.String(), "http://www.alnago.com/index.php/frontend/ajax/es/") {
			pagenum=pagenum+8
			arr := []byte("order=id+DESC&view=grid&page_num="+ strconv.Itoa(pagenum) +"&v_search_option_4=Arriendo")
			//fmt.Println("PostRaw: " + string(arr))
			c.PostRaw("http://www.alnago.com/index.php/frontend/ajax/es/1/"+ strconv.Itoa(pagenum), arr)
		
		} else if strings.HasPrefix(r.Request.URL.String(), "http://www.alnago.com/index.php/property") {
			ScrapedAsset.Code=strings.Split(r.Request.URL.String(), "/")[5]
			ScrapedAsset.Status=true
			ScrapedAsset.Agency=agencyname
			ScrapedAsset.Link = r.Request.URL.String()
			db.DBInsertRedis(ScrapedAsset.GetCode(),ScrapedAsset.ToJSON())
		}
	})


	// On every a element which has href attribute call callback
	c.OnHTML("p.bottom-border", func(e *colly.HTMLElement) {
		TextTitle := strings.TrimSpace(e.ChildText("strong"))
		switch TextTitle {
			case "Dirección": ScrapedAsset.Sector=e.ChildText("span")
			case "Proposito:": ScrapedAsset.Business=e.ChildText("span")
			case "Precio de Arriendo:": ScrapedAsset.Price=e.ChildText("span")
			case "Tipo:" :ScrapedAsset.Type=e.ChildText("span")
			case "Tamaño Exacto:":ScrapedAsset.Area=e.ChildText("span")
			case "Alcobas:": ScrapedAsset.Numrooms=e.ChildText("span")
			case "Baños:": ScrapedAsset.Numbaths=e.ChildText("span")
			}


	})
	c.OnResponse(func(r *colly.Response) {
		rs := &Responsejson{}
		_ = json.Unmarshal(r.Body, rs)
		//fmt.Println(rs.Print)
			
			doc, _ := goquery.NewDocumentFromReader(bytes.NewReader([]byte(rs.Print)))
			doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
				// For each item found, get the band and title
				link,_ := s.Attr("href")
				if strings.HasPrefix(link, "http://www.alnago.com/index.php/property") {
					fmt.Println("LINK>> "+link) 
					//c.Visit(link)
				}
				
			  })
	})
	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		arr := []byte("order=id+DESC&view=grid&page_num="+ strconv.Itoa(pagenum) +"&v_search_option_4=Arriendo")
		fmt.Println("RETRY")
		c.PostRaw("http://www.alnago.com/index.php/frontend/ajax/es/1/"+ strconv.Itoa(pagenum), arr)
	
	})
	arr := []byte("order=id+DESC&view=list&page_num="+ strconv.Itoa(pagenum) +"&v_search_option_4=Arriendo")
	c.PostRaw("http://www.alnago.com/index.php/frontend/ajax/es/1/" + strconv.Itoa(pagenum), arr)

	fmt.Println("Collect end at: " + time.Now().Format(time.Stamp) )
}

