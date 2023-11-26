package main

import(
    "fmt"
    "log"
    "github.com/gocolly/colly"
)

func main(){
    c := colly.NewCollector()

    c.OnRequest(func(r *colly.Request) {
	fmt.Println("Visiting: ", r.URL)
    })

    c.OnError(func(_ *colly.Response, err error) {
	log.Println("Something went wrong: ", err)
    })

    c.OnResponse(func(r *colly.Response) {
	fmt.Println("Page visited: ", r.Request.URL)
    })

    c.OnHTML("a", func(e *colly.HTMLElement) {
	fmt.Printf("%v\n", e.Attr("href"))
    })

    c.OnScraped(func(r *colly.Response) {
	fmt.Println(r.Request.URL, " scraped!")
    })


    c.Visit("https://scrapeme.live/shop/")
}
