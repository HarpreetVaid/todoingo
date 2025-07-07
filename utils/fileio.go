/*
Copyright © 2025 Harpreet Singh harry798677@gmail.com
*/
package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func Getfilecontent(cfgFile string) string {
	if fileExists(cfgFile) {
		data, err := os.ReadFile(cfgFile)
		if err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}
		return string(data)
	} else {
		// Get current working directory
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			os.Exit(1)
		}

		// Join with the given relative path
		fullPath := filepath.Join(cwd, cfgFile)

		// Create the directory structure if it doesn't exist
		err = os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			os.Exit(1)
		}

		// Create the file at the full path
		file, err := os.Create(fullPath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			os.Exit(1)
		}
		defer file.Close()

		// Write empty JSON array to file
		_, err = file.WriteString("[]")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}

		return "[]"
	}
}

func Savefilecontent(userData string, cfgFile string) bool {
	content := []byte(userData)
	err := os.WriteFile(cfgFile, content, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return false
	}
	fmt.Println("File Upadated successfully.")
	return true
}
