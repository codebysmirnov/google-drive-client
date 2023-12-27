package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"google-drive-client/client"
)

func readFile(filename string) ([]byte, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return fileContent, err
}

func main() {
	c, err := client.NewClient("credentials.json", []string{"1sIWkMtkpoeDlRw_jsz0eC5PjqtQC3ExG"})
	if err != nil {
		log.Fatal(err)
	}

	filename := "example.txt"
	fileContent, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	contentType := http.DetectContentType(fileContent)
	fileID, err := c.UploadFile(filename, contentType, fileContent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", fileID)
}
