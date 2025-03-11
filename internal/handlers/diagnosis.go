package handlers

import (
        "encoding/json"
        "net/http"
        "os"
        "os/exec"
        "path/filepath"

        "github.com/gin-gonic/gin"
)

func DiagnosisHandler(c *gin.Context) {
        uuid := c.Param("UUID")

        filePathLR := filepath.Join("images/lr", uuid)
        filePathSR := filepath.Join("images/sr", uuid)

        var filePath string
        if _, err := os.Stat(filePathSR); err == nil {
                filePath = filePathSR
        } else if _, err := os.Stat(filePathLR); err == nil {
                filePath = filePathLR
        } else {
                c.JSON(http.StatusNotFound, gin.H{"error": "File không tồn tại"})
                return
        }

        cmd := exec.Command("python3", "scripts/info.py", filePath)
        output, err := cmd.Output()
        if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi chạy script Python"})
                return
        }

        var result map[string]interface{}
        err = json.Unmarshal(output, &result)
        if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi parse JSON từ Python"})
                return
        }

        c.JSON(http.StatusOK, result)
}
