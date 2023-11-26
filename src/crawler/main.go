// package main

// import (
// 	"fmt"

// 	"github.com/gocolly/colly"
// )

// func main() {
// 	// Instantiate default collector
// 	c := colly.NewCollector(
// 		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
// 		// colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
// 		colly.AllowedDomains("news.ycombinator.com"),
// 	)

// 	// On every a element which has href attribute call callback
// 	c.OnHTML("span.titleline > a[href]", func(e *colly.HTMLElement) {
// 		link := e.Attr("href")
// 		// Print link
// 		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
// 		// Visit link found on page
// 		// Only those links are visited which are in AllowedDomains
// 		c.Visit(e.Request.AbsoluteURL(link))
// 	})

// 	// Before making a request print "Visiting ..."
// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL.String())
// 	})

// 	// Start scraping on https://news.ycombinator.com
// 	c.Visit("https://news.ycombinator.com/")
// }

// Command screenshot is a chromedp example demonstrating how to take a
// screenshot of a specific element and of the entire browser viewport.
package main

import (
	// "context"
	// "log"
	// "os"
	"fmt"
	"time"

	// "github.com/chromedp/chromedp"
	// "github.com/academy-software/go-aws-sqs-consumer"
	consumer "github.com/academy-software/go-aws-sqs-consumer"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	q := "https://sqs.us-east-1.amazonaws.com/267074127319/test"

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           "leonardo",
	}))

	c := consumer.New(q, handle,
		&consumer.Config{
			AwsSession:                  sess,
			Receivers:                   1,
			SqsMaxNumberOfMessages:      10,
			SqsMessageVisibilityTimeout: 20,
			PollDelayInMilliseconds:     100,
		})

	c.Start()
}

func handle(m *sqs.Message) error {
	fmt.Println("Message Body:", *(m.Body))
	//emulate processing time
	time.Sleep(time.Second * 2)
	return nil
}

// func handler() {
// 	// create context
// 	ctx, cancel := chromedp.NewContext(
// 		context.Background(),
// 		// chromedp.WithDebugf(log.Printf),
// 	)
// 	defer cancel()

// 	// capture screenshot of an element
// 	var buf []byte
// 	if err := chromedp.Run(ctx, elementScreenshot(`https://pkg.go.dev/`, `img.Homepage-logo`, &buf)); err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := os.WriteFile("elementScreenshot.png", buf, 0o644); err != nil {
// 		log.Fatal(err)
// 	}

// 	// capture entire browser viewport, returning png with quality=90
// 	if err := chromedp.Run(ctx, fullScreenshot(`https://brank.as/`, 90, &buf)); err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := os.WriteFile("fullScreenshot.png", buf, 0o644); err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Printf("wrote elementScreenshot.png and fullScreenshot.png")
// }

// // elementScreenshot takes a screenshot of a specific element.
// func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
// 	return chromedp.Tasks{
// 		chromedp.Navigate(urlstr),
// 		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
// 	}
// }

// // fullScreenshot takes a screenshot of the entire browser viewport.
// //
// // Note: chromedp.FullScreenshot overrides the device's emulation settings. Use
// // device.Reset to reset the emulation and viewport settings.
// func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
// 	return chromedp.Tasks{
// 		chromedp.Navigate(urlstr),
// 		chromedp.FullScreenshot(res, quality),
// 	}
// }
