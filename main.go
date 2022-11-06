package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.StaticFS("/", gin.Dir(".", true))

	currDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("get current directory path failed: %s", err.Error())
	}

	log.Printf("serving %s on :8080", currDir)
	router.Run(":8080")
}
