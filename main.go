package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

// func Quote(context *gin.Context)  {
// 	context.IndentedJSON(http.StatusOK, responseData)
// }

func Quote() string {
	response, err := http.Get("https://api.kanye.rest")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp := string (responseData)
	ourNames := resp + " Amanda & Kalese"
	return ourNames

	// fmt.Printf(string(responseData) + " " + "Amanda&Kalese")

	// router := gin.Default() 
	// router.Run("localhost:9090")

}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": Quote(),
	  })
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
  }
