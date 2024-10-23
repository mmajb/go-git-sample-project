package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mmajb/go-git-sample-project/math"
	"github.com/mmajb/go-git-sample-project/utils"
)

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
}
