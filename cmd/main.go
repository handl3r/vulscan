package main

import (
	"vulscan/api/http"
	"vulscan/bootstrap"
	"vulscan/configs"
)

func main() {
	configs.LoadConfig()
	appContext := bootstrap.LoadServices(configs.Get())

	controllerManager := bootstrap.LoadControllerManager(appContext)
	server := http.NewServer(http.NewRouter(appContext, controllerManager))
	server.Run(configs.Get().ServerAddress)
}

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

//package main
//
//import (
//	"fmt"
//
//	"github.com/gocolly/colly"
//)
//
//func main() {
//	// Instantiate default collector
//	c := colly.NewCollector(
//		colly.AllowedDomains("juice-shop.herokuapp.com"),
//		colly.MaxDepth(2),
//	)
//
//	// On every a element which has href attribute call callback
//	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
//		link := e.Attr("href")
//		// Print link
//		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
//		// Visit link found on page
//		// Only those links are visited which are in AllowedDomains
//		c.Visit(e.Request.AbsoluteURL(link))
//	})
//
//	// Before making a request print "Visiting ..."
//	c.OnRequest(func(r *colly.Request) {
//		fmt.Println("Visiting", r.URL.String())
//	})
//
//	c.Visit("http://juice-shop.herokuapp.com/#/")
//}

//package main

//
//func main() {
//
//	s := "postgres://user:pass@host.com:40/path?k=v,1&t=2#f"
//
//	u, err := url.Parse(s)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println("scheme: ", u.Scheme)
//
//	fmt.Println("user: ", u.User)
//	fmt.Println("username: ", u.User.Username())
//	p, _ := u.User.Password()
//	fmt.Println("password: ", p)
//
//	fmt.Println("host: ", u.Host)
//	host, port, _ := net.SplitHostPort(u.Host)
//	fmt.Println("host: ", host)
//	fmt.Println("port: ", port)
//
//	fmt.Println("path", u.Path)
//	fmt.Println(u.Fragment)
//
//	fmt.Println(u.RawQuery)
//	m, i := url.ParseQuery(u.RawQuery)
//	fmt.Println(m, i)
//	fmt.Println(m["k"][0])
//}
