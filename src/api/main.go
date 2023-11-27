package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gin-gonic/gin"
)

// products represents data about a record album.
type Product struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// products slice to seed record album data.
var products = []Product{
	{Title: "Jumbo Yoghurt Griekse Stijl  Naturel 10% Vet 1kg", URL: "https://jumbo.com/dam-images/fit-in/360x360/Products/18092023_1695036579191_1695036591485_8718452394951_1.png"},
	{Title: "Jumbo Champignons Voordeelverpakking 400g", URL: "https://jumbo.com/dam-images/fit-in/360x360/Products/29092023_1695996129860_1695996141161_8718452601240_1.png"},
}

func main() {
	r := gin.Default()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// Define your routes here
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/products", getProducts)

	r.POST("/products", postProducts)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// // getProducts responds with the list of all albums as JSON.
func getProducts(c *gin.Context) {

	bucket := os.Getenv("S3_BUCKET")

	if bucket == "" {
		log.Fatal("S3_BUCKET not found")
	}

	s3Url := os.Getenv("S3_URL")

	if bucket == "" {
		log.Fatal("S3_URL not found")
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create an S3 client
	svc := s3.New(sess)

	// Create an input object
	input := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	}

	// // List the objects in the bucket
	result, err := svc.ListObjects(input)
	if err != nil {
		fmt.Println("Error listing objects:", err)
		return
	}

	var objectKeys []string
	// // Print the objects
	for _, object := range result.Contents {
		fmt.Println(*object.Key)
		objectKeys = append(objectKeys, s3Url+"/"+*object.Key)
	}

	c.IndentedJSON(http.StatusOK, objectKeys)
}

type response struct {
	Endpoint string
}

// postAlbums adds an album from JSON received in the request body.
func postProducts(c *gin.Context) {

	var newProducts []Product

	// Call BindJSON to bind the received JSON to
	// newProduct.
	if err := c.BindJSON(&newProducts); err != nil {
		return
	}

	// TODO put that as var
	q := os.Getenv("SQS_QUEUE_URL")

	println("queue: ", q)

	if q == "" {
		log.Fatal("SQS_QUEUE_URL not found")
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Marshal JSON data into a byte slice
	jsonData, err := json.Marshal(newProducts)
	if err != nil {
		fmt.Println("Error marshaling JSON data:", err)
		return
	}

	// Convert byte slice to a string
	messageBody := string(jsonData)
	fmt.Println("JSON string:", messageBody)

	// messageBody := newProducts[].(string)
	err = SendMessage(sess, q, messageBody)
	if err != nil {
		fmt.Printf("Got an error while trying to send message to queue: %v", err)
		return
	}

	fmt.Println("Message sent successfully")
	// Add the new album to the slice.
	products = append(products, newProducts...)
	c.IndentedJSON(http.StatusCreated, newProducts)
}

func SendMessage(sess *session.Session, queueUrl string, messageBody string) error {
	sqsClient := sqs.New(sess)

	_, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    &queueUrl,
		MessageBody: aws.String(messageBody),
	})

	return err
}
