package main

import (
        "github.com/gin-gonic/gin"
        "su-api/internal/handlers"
)

func main() {
        r := gin.Default()

        r.POST("/upload", handlers.UploadHandler)
        r.GET("/download/:filename", handlers.DownloadHandler)
        r.GET("/diagnosis/:UUID", handlers.DiagnosisHandler)

        r.Run(":8080")
}
