package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"text.com/initializer"
	"text.com/models"
	"text.com/routers"
)

func init() {
	initializer.LoadEnvVeriables()
	initializer.RunOnPort()

	// Connect To to and store the instance
	initializer.ConnectToDb()
	initializer.SyscDb()

}

func main() {

	log.Print("\n\nReady to Start The Server --> \n\n")

	// Setup Server
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	// Routes
	// test Route
	r.GET("/ping", pong)

	// Attached Routes
	routers.UserRoutes(r)
	routers.LoanRouters(r)

	// error print
	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}

}

// Sample Surver Status Cheaking function
func pong(c *gin.Context) {
	c.JSON(http.StatusOK, models.SendMessage{
		StatusCode: http.StatusOK,
		Message:    "pong , Server Is On",
	})
}
