package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// type kanyeSays struct {
// 	Quote string `json:"quote"`
// }

// Get Quote from API and return as string
func Quote() string {
	response, err := http.Get("https://api.kanye.rest")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// Mapping response body
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Turn response to a string from a Byte + concatenate our names
	resp := string(responseData)
	ourNames := resp + " Amanda & Kalese"
	return ourNames
}

// Starting the server on PS route
func main() {
	r := gin.Default()
	r.GET("/ps", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": Quote(),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
