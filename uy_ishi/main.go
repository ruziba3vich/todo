package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	sourceFileName := "uy_ishi/abduazim.txt"

	sourceFile, err := os.OpenFile(sourceFileName, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer sourceFile.Close()

	backupFileName := generateBackupFileName(sourceFileName)

	backupFile, err := os.Create(backupFileName)
	if err != nil {
		fmt.Println("Error creating backup file:", err)
		return
	}
	defer backupFile.Close()

	_, err = io.Copy(backupFile, sourceFile)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}

	fmt.Printf("Backup file successfully created: %s\n", backupFileName)
}

func generateBackupFileName(sourceFileName string) string {
	currentTime := time.Now()

	dateString := currentTime.Format("2006-01-02")

	timeString := currentTime.Format("150405")

	extension := filepath.Ext(sourceFileName)

	backupFileName := fmt.Sprintf("%s_backup_%s_%s%s", sourceFileName[:len(sourceFileName)-len(extension)], dateString, timeString, extension)

	return backupFileName
}
