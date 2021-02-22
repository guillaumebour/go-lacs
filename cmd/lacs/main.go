package main

import (
	"LaTeXCompilationService/internal/server"
	"fmt"
)

func main() {
	lacsOptions := server.LacsOptions{
		UploadFileSizeLimit: 8 << 20, // 8 MiB
	}
	router := server.NewLacsRouter(lacsOptions)

	if err := router.Run(":8081"); err != nil {
		fmt.Printf("error while serving: %v", err)
	}
}
