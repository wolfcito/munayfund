package fileinfra

import (
	"path/filepath"
	"strings"
)

func GetTypeByExtension(filePath string) string {
	switch strings.ToLower(filepath.Ext(filePath)) {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp":
		return "image"
	case ".mp4", ".avi", ".mkv", ".mov":
		return "video"
	default:
		return "otro"
	}
}
