//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	fmt.Println("Hello")
//}
//
////package main
//
////func main() {
////	// create context
////	ctx, cancel := chromedp.NewContext(context.Background())
////	defer cancel()
////
////	// run task list
////	var nodes []*cdp.Node
////	err := chromedp.Run(ctx,
////		chromedp.Navigate(`https://app.diagrams.net/`),
////		chromedp.Nodes("img", &nodes),
////	)
////	if err != nil {
////		log.Fatal(err)
////	}
////
////	for _, node := range nodes {
////		fmt.Println("x", node.AttributeValue("src"))
////	}
////}

package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("handl3r.netlify.app"),
		colly.MaxDepth(1),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://handl3r.netlify.app")
}