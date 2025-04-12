// Package lassie provides a client for interacting with Lassie - an IPFS retrieval client.
// Lassie specializes in fetching content from IPFS using various retrieval protocols.
//
// Key Features:
// - Simple HTTP interface for IPFS content retrieval
// - Built-in timeout handling
// - Support for both raw data and JSON content
//
// See https://github.com/filecoin-project/lassie for more details.
package lassie

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// DefaultTimeout specifies the maximum duration for HTTP requests to Lassie
const DefaultTimeout = 20 * time.Second

// LassieClient represents a client for interacting with a Lassie IPFS retrieval node.
// It handles communication with the node and content retrieval operations.
type LassieClient struct {
	// Scheme is the protocol scheme (e.g. "http", "https") used for communication
	scheme string

	// Host is the hostname or IP address of the Lassie node (e.g. "localhost")
	host string

	// Port is the TCP port number used for communication (e.g. 8080)
	port int

	// LassieClientURL is the base URL for IPFS content retrieval
	// Format: "{scheme}://{host}:{port}/ipfs"
	LassieClientURL string
}

// NewLassieClient creates a new LassieClient instance configured to communicate
// with a Lassie node at the specified location.
//
// Example:
//   client := NewLassieClient("http", "localhost", 8080)
//   data, err := client.GetDataFromCID("bafy...")
func NewLassieClient(scheme string, host string, port int) *LassieClient {

	return &LassieClient{
		scheme:          scheme,
		host:            host,
		port:            port,
		LassieClientURL: fmt.Sprintf("%s://%s:%d/ipfs", scheme, host, port),
	}

}

// GetURLFromCID returns the URL of the given CID on the Lassie node.
//
// The function constructs the URL by appending the provided CID to the LassieClientURL of the LassieClient.
//
// Parameters:
//   - cid (string) - The content identifier (CID) of the data to retrieve.
//
// Returns:
//   - cidUrl (string) - The constructed URL of the given CID.
//   - err (error) - An error if any occurs while constructing the URL.
func (lassieC *LassieClient) GetURLFromCID(cid string) (cidUrl string, err error) {
	cidUrl = fmt.Sprintf("%s/%s", lassieC.LassieClientURL, cid)

	return cidUrl, err
}

// GetDataFromCID retrieves IPFS content by its CID from the Lassie node.
// The content is returned as raw bytes which can represent any IPFS content type.
//
// Parameters:
//   - cid: The content identifier (CID) of the data to retrieve (e.g. "bafy...")
//
// Returns:
//   - []byte: The raw content data
//   - error: Any error that occurred during retrieval, including:
//     - Invalid CID format
//     - Connection failures
//     - Timeouts (after 20 seconds)
//     - HTTP errors (non-200 status)
func (lassieC *LassieClient) GetDataFromCID(cid string) ([]byte, error) {
	cidUrl := fmt.Sprintf("%s/%s", lassieC.LassieClientURL, cid)

	req, err := http.NewRequest("GET", cidUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "*/*")

	client := &http.Client{Timeout: DefaultTimeout}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d retrieving CID %s", resp.StatusCode, cid)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return data, nil
}

// GetJSONFromCID retrieves the JSON data associated with the given CID from the Lassie node.
//
// Parameters:
//   - cid (string) - The content identifier (CID) of the JSON data to retrieve.
//
// Returns:
//   - jsonStr (string) - The retrieved JSON data as a string.
//   - err (error) - An error if any occurs while retrieving the data.
func (lassieC *LassieClient) GetJSONFromCID(cid string) (jsonStr string, err error) {
	// Retrieve the data from the CID.
	data, err := lassieC.GetDataFromCID(cid)
	if err != nil {
		return "", err
	}

	// Convert the data to a string.
	jsonStr = string(data)
	return jsonStr, nil
}
