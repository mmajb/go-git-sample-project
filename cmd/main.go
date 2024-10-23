package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mmajb/go-git-sample-project/math"
	"github.com/mmajb/go-git-sample-project/utils"
	"github.com/gin-gonic/gin"
    //"github.com/swaggo/swag/example/basic/docs" // you will need to import this for the generated docs
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files" // swagger embed files

)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

func main() {
	fmt.Println("Go Sample Project on github")

	// Using the math package
	sum := math.Add(5, 10)
	fmt.Printf("Sum of 5 and 10: %d\n", sum)

	// Using the utils package
	greeting := utils.Greet("Alice")
	fmt.Println(greeting)

	// Using the external github.com/google/uuid package
	id := uuid.New()
	fmt.Printf("Generated UUID: %s\n", id.String())
	
	r := gin.Default()

    // Swagger route
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Simple API endpoint
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    // Start the server
    r.Run(":8080")
}
