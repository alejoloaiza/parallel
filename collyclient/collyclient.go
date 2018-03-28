package collyclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
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
		li := e.Attr("li")
		//fmt.Println("Li Found: " + e.Text + ":" + li)
		if strings.Contains(e.Text,"Sector") {
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
		}
	})


	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		
		if strings.HasPrefix(link, "/webs/santafe/Pages/recomendar") {
			return
		}
		if strings.Contains(link,".pdf")  {
			return
		}
		c.Visit(e.Request.AbsoluteURL(link))
	})



	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("http://www.arrendamientossantafe.com/webs/santafe/pages/basico?bussines_type=Arrendar")
}




func generateFormData() map[string][]byte {
	f, _ := os.Open("gocolly.png")
	defer f.Close()

	imgData, _ := ioutil.ReadAll(f)

	return map[string][]byte{
		"firstname": []byte("one"),
		"lastname":  []byte("two"),
		"email":     []byte("onetwo@example.com"),
		"file":      imgData,
	}
}

func setupServer() {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("received request")
		err := r.ParseMultipartForm(10000000)
		if err != nil {
			fmt.Println("server: Error")
			w.WriteHeader(500)
			w.Write([]byte("<html><body>Internal Server Error</body></html>"))
			return
		}
		w.WriteHeader(200)
		fmt.Println("server: OK")
		w.Write([]byte("<html><body>Success</body></html>"))
	}

	go http.ListenAndServe(":8080", handler)
}

func Multipart() {
	// Start a single route http server to post an image to.
	setupServer()

	c := colly.NewCollector(colly.AllowURLRevisit(), colly.MaxDepth(5))

	// On every a element which has href attribute call callback
	c.OnHTML("html", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		time.Sleep(1 * time.Second)
		e.Request.PostMultipart("http://localhost:8080/", generateFormData())
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Posting gocolly.jpg to", r.URL.String())
	})

	// Start scraping
	c.PostMultipart("http://localhost:8080/", generateFormData())
	c.Wait()
}

