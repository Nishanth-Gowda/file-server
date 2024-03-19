package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nishanth-gowda/file-server/internal/filestore"
	"github.com/nishanth-gowda/file-server/internal/repository"
)

func main() {
	repo, err := repository.NewFileRepository("./storage")
	if err != nil {
		log.Fatalf("Failed to create file repository: %v", err)
	}

	fs := filestore.NewFileStore(repo)

	router := gin.Default()

	// Register routes
	router.POST("/store", fs.HandleStore)
	router.PUT("/update", fs.HandleUpdate)
	router.DELETE("/delete", fs.HandleDelete)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
