package main 
 
import ( 
	"encoding/csv" 
	"github.com/gocolly/colly" 
	"log" 
	"os" 
) 
 
type PokemonProduct struct { 
	url, image, name, price string 
} 
 
func contains(s []string, str string) bool { 
	for _, v := range s { 
		if v == str { 
			return true 
		} 
	} 
 
	return false 
} 
 
func main() { 
	var pokemonProducts []PokemonProduct 
 
	var pagesToScrape []string 
 
	pageToScrape := "https://scrapeme.live/shop/page/1/" 
 
	pagesDiscovered := []string{ pageToScrape } 
 
	i := 1 
	limit := 5 
 
	c := colly.NewCollector() 
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36" 
 
	c.OnHTML("a.page-numbers", func(e *colly.HTMLElement) { 
		newPaginationLink := e.Attr("href") 
 
		if !contains(pagesToScrape, newPaginationLink) { 
			if !contains(pagesDiscovered, newPaginationLink) { 
				pagesToScrape = append(pagesToScrape, newPaginationLink) 
			} 
			pagesDiscovered = append(pagesDiscovered, newPaginationLink) 
		} 
	}) 
 
	c.OnHTML("li.product", func(e *colly.HTMLElement) { 
		pokemonProduct := PokemonProduct{} 
 
		pokemonProduct.url = e.ChildAttr("a", "href") 
		pokemonProduct.image = e.ChildAttr("img", "src") 
		pokemonProduct.name = e.ChildText("h2") 
		pokemonProduct.price = e.ChildText(".price") 
 
		pokemonProducts = append(pokemonProducts, pokemonProduct) 
	}) 
 
	c.OnScraped(func(response *colly.Response) { 
		if len(pagesToScrape) != 0 && i < limit { 
			pageToScrape = pagesToScrape[0] 
			pagesToScrape = pagesToScrape[1:] 
 
			i++ 
 
			c.Visit(pageToScrape) 
		} 
	}) 
 
	c.Visit(pageToScrape) 
 
	file, err := os.Create("products.csv") 
	if err != nil { 
		log.Fatalln("Failed to create output CSV file", err) 
	} 
	defer file.Close() 
 
	writer := csv.NewWriter(file) 
 
	headers := []string{ 
		"url", 
		"image", 
		"name", 
		"price", 
	} 
	writer.Write(headers) 
 
	for _, pokemonProduct := range pokemonProducts { 
		record := []string{ 
			pokemonProduct.url, 
			pokemonProduct.image, 
			pokemonProduct.name, 
			pokemonProduct.price, 
		} 
 
		writer.Write(record) 
	} 
	defer writer.Flush() 
}
