package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"sync"

	consumer "github.com/academy-software/go-aws-sqs-consumer"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Product struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

var bucket = ""

func main() {
	// TODO put that as var
	q := os.Getenv("SQS_QUEUE_URL")

	if q == "" {
		log.Fatal("SQS_QUEUE_URL not found")
	}

	bucket = os.Getenv("S3_BUCKET")

	if bucket == "" {
		log.Fatal("S3_BUCKET not found")
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		// TODO put that as var
		// Profile: "leonardo",
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

func handle(message *sqs.Message) error {

	products := []Product{}
	if err := json.Unmarshal([]byte(*message.Body), &products); err != nil {
		log.Fatal("Error unmarshaling message body:", err)
	}

	fmt.Println("Products:", products[0].URL)

	for _, info := range products {
		applyWaterMark(info.Title, info.URL)
	}

	fmt.Println("Screenshot Taken")

	return nil
}

func applyWaterMark(title string, imageUrl string) {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithErrorf(log.Printf))
	defer cancel()

	// do this first so that its page.EventLoadEventFired event won't be caught
	if err := chromedp.Run(ctx,
		chromedp.Navigate("about:blank"),
	); err != nil {
		log.Fatal(err)
	}

	imageFilename := extractImageNameFromURL(imageUrl)

	var wg sync.WaitGroup
	wg.Add(1)
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev.(type) {
		case *page.EventLoadEventFired:
			go func() {
				var data []byte
				if err := chromedp.Run(ctx,
					chromedp.Screenshot("div.parent > .image", &data, chromedp.NodeVisible),
				); err != nil {
					fmt.Println("fail")
				}

				if err := uploadS3("with_banner_"+imageFilename, data);
				// if err := os.WriteFile("screenshots/with_banner_"+imageFilename, data, 0644);
				err != nil {
					fmt.Println("upload s3 failed")
				}
				wg.Done()
			}()
		}
	})

	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	  <meta charset="UTF-8">
	  <meta name="viewport" content="width=device-width, initial-scale=1.0">
	  <style>
		.image {
		  width: 100%;
		  height: 100%;
		  
		}
	
	  .parent {
		overflow: hidden; /* required */

		margin: 25px auto; /* for demo only */
		border:1px solid grey; /* for demo only */
		position: relative; /* required  for demo*/
	  }
	  
	  .ribbon {
		margin: 0;
		padding: 0;
		background: red;
		color:white;
		text-shadow: 1px 1px 1px black;
		padding: 2em 0;
		position: absolute;
		top:0;
		right:0;
		transform: translateX(30%) translateY(0%) rotate(45deg);
		transform-origin: top left;
	  }
	  .ribbon:before,
	  .ribbon:after {
		content: '';
		position: absolute;
		top:0;
		margin: 0 -1px; /* tweak */
		width: 100%;
		height: 100%;
		background: red;
	  }
	  .ribbon:before {
		right:100%;
	  }
	
	  .ribbon:after {
		left:100%;
	  }

	  .overlay {
		position: absolute;
		left: 0;
		bottom: 0;
		width: 100%;
		height: 10%;
		margin-bottom: 20px;
		background-color: #eeb717;
		opacity: 0.9;
		display: flex;
		justify-content: center;
		align-items: center;
		color: black;
		font-size: 18px;
		font-weight: bold;
	  }
	  </style>
	</head>
	<body>
	<div class="parent">
	  <img class="image" src="` + imageUrl + `" alt="Black Friday Image">
		<h4 class="ribbon">Special Sale Today</h4>
		<div class="overlay">` + title + `</div>
	  </div>
	</body>
	</html>
	`
	if err := chromedp.Run(ctx,
		chromedp.PollFunction("(html) => {document.open();document.write(html);document.close();return true;}", nil, chromedp.WithPollingArgs(html)),
	); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}

func extractImageNameFromURL(imageUrl string) string {

	// Parse the URL
	parsedURL, err := url.Parse(imageUrl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
	}

	// Extract the image name from the URL path
	imageName := path.Base(parsedURL.Path)

	fmt.Println("Image Name:", imageName)

	return imageName
}

func uploadS3(filename string, data []byte) error {

	sesss := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		// TODO put that as var
		// Profile: "leonardo",
	}))

	// file, header, err := r.FormFile("file")
	// if err != nil {
	// 	// Do your error handling here
	// 	return
	// }
	// defer file.Close()

	// filename := header.Filename

	uploader := s3manager.NewUploader(sesss)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucket),      // Bucket to be used
		Key:         aws.String(filename),    // Name of the file to be saved
		Body:        bytes.NewReader(data),   // File
		ContentType: aws.String("image/png"), // content type
		ACL:         aws.String("public-read"),
	})

	if err != nil {
		// Do your error handling here
		return err
	}

	return nil
}
