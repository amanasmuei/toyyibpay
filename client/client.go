package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

// Client represents a ToyyibPay API client.
type Client struct {
	APIKey  string
	BaseURL string
}

// NewClient initializes a new ToyyibPay client.
func NewClient() *Client {
	return &Client{
		APIKey:  os.Getenv("TOYYIBPAY_API_KEY"),
		BaseURL: os.Getenv("TOYYIBPAY_BASE_URL"),
	}
}

// CreateBill sends a request to create a new bill in ToyyibPay.
func (c *Client) CreateBill(req interface{}) (*http.Response, error) {
	url := c.BaseURL + "/index.php/api/createBill"
	jsonData, err := json.Marshal(req)
	if err != nil {
		log.WithError(err).Error("Failed to marshal request")
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.WithError(err).Error("Failed to send request")
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, &APIError{StatusCode: resp.StatusCode, Message: "Error creating bill"}
	}

	return resp, nil
}
