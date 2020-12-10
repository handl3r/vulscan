package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
}

//package main

//func main() {
//	// create context
//	ctx, cancel := chromedp.NewContext(context.Background())
//	defer cancel()
//
//	// run task list
//	var nodes []*cdp.Node
//	err := chromedp.Run(ctx,
//		chromedp.Navigate(`https://app.diagrams.net/`),
//		chromedp.Nodes("img", &nodes),
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, node := range nodes {
//		fmt.Println("x", node.AttributeValue("src"))
//	}
//}