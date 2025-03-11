package handlers

import (
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func UploadHandler(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Không thể đọc file"})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".dcm" && ext != ".dicom" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Chỉ hỗ trợ DICOM (.dcm, .dicom)"})
		return
	}

	uuidStr := uuid.New().String()
	lrPath := filepath.Join("images/lr", uuidStr+ext)

	os.MkdirAll("images/lr", os.ModePerm)
	c.SaveUploadedFile(file, lrPath)

	// Chỉ xử lý DICOM, bỏ PNG
	cmd := exec.Command("python3", "scripts/process.py", "-i", lrPath, "-o", "images/sr/"+uuidStr+".dcm")
	err = cmd.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi xử lý ảnh DICOM"})
		return
	}

	response := gin.H{
		"message": "Upload thành công",
		"lr_path": lrPath,
		"sr_dicom": "images/sr/" + uuidStr + ".dcm",
	}

	c.JSON(http.StatusOK, response)
}
