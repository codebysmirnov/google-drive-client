package client

import (
	"bytes"
	"fmt"

	"google.golang.org/api/drive/v3"
)

// UploadFile uploads a file to Google Drive using the provided filename, contentType, and fileContent.
// It returns the ID of the newly uploaded file or an error if the upload fails.
func (c *Client) UploadFile(filename, contentType string, fileContent []byte) (string, error) {
	// Create a drive.File object with the specified metadata.
	f := &drive.File{
		Name:     filename,
		Parents:  c.parents,
		MimeType: contentType,
	}

	// Upload the file content using the Google Drive service's Files.Create method.
	res, err := c.service.Files.Create(f).Media(bytes.NewReader(fileContent)).Do()
	if err != nil {
		return "", fmt.Errorf("can't create file: %w", err)
	}

	// Return the ID of the newly uploaded file.
	return res.Id, nil
}
