// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()

// 	// Configure Gin to trust proxy headers
// 	router.Use(gin.Recovery())
// 	router.Use(gin.Logger())

// 	// Define your routes here
// 	router.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
// 	})

// 	router.GET("/products", getProducts)
// 	router.POST("/products", postProducts)

// 	router.Run("localhost:8080")
// }

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// products represents data about a record album.
type product struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// products slice to seed record album data.
var products = []product{
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
	c.IndentedJSON(http.StatusOK, products)
}

// postAlbums adds an album from JSON received in the request body.
func postProducts(c *gin.Context) {
	// TODO put that as var
	q := os.Getenv("SQS_QUEUE_URL")

	if q == "" {
		log.Fatal("SQS_QUEUE_URL not found")
	}

	var newProducts []product

	// Call BindJSON to bind the received JSON to
	// newProduct.
	if err := c.BindJSON(&newProducts); err != nil {
		return
	}

	// Add the new album to the slice.
	products = append(products, newProducts...)
	c.IndentedJSON(http.StatusCreated, newProducts)
}
