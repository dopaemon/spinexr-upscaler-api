package handlers

import (
        "github.com/gin-gonic/gin"
        "net/http"
        "os"
        "path/filepath"
)

func DownloadHandler(c *gin.Context) {
        filename := c.Param("filename")
        filePath := filepath.Join("images/sr", filename)

        if _, err := os.Stat(filePath); os.IsNotExist(err) {
                c.JSON(http.StatusNotFound, gin.H{"error": "File không tồn tại"})
                return
        }

        c.File(filePath)
}
