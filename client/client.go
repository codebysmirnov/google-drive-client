package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

// Client represents a client for interacting with the Google Drive API.
type Client struct {
	service *drive.Service // Google Drive API service.
	parents []string       // Parent folder IDs for uploading files.
}

// NewClient creates a new instance of Client for interacting with the Google Drive API.
func NewClient(credentialsFilePath string, parents []string) (*Client, error) {
	// Load credentials from a file.
	c, err := loadCredentialsFromFile(credentialsFilePath)
	if err != nil {
		return nil, err
	}

	// Create an HTTP client for authentication.
	client := newClient(c)

	// Create the Google Drive service using the HTTP client.
	s, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("can't construct client: %w", err)
	}

	// Initialize and return a Client instance.
	return &Client{
		service: s,
		parents: parents,
	}, nil
}

// credentials represents the data structure for parsing credentials from a JSON file.
type credentials struct {
	Email      string `json:"client_email"`
	PrivateKey string `json:"private_key"`
}

// loadCredentialsFromFile loads credentials from a JSON file.
func loadCredentialsFromFile(credentialFile string) (credentials, error) {
	// Read the contents of the file.
	b, err := os.ReadFile(credentialFile)
	if err != nil {
		return credentials{}, fmt.Errorf("read credentials file: %w", err)
	}

	// Unmarshal credentials from JSON.
	var c credentials
	err = json.Unmarshal(b, &c)
	if err != nil {
		return credentials{}, fmt.Errorf("unmarshal credentials: %w", err)
	}

	// Return the credentials structure.
	return c, nil
}

// newClient creates and returns an HTTP client for authentication using the credentials.
func newClient(c credentials) *http.Client {
	// JWT configuration to create the HTTP client.
	config := &jwt.Config{
		Email:      c.Email,
		PrivateKey: []byte(c.PrivateKey),
		Scopes: []string{
			drive.DriveScope,
		},
		TokenURL: google.JWTTokenURL,
	}

	// Create and return the HTTP client.
	return config.Client(context.TODO())
}
