package lassie

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// LassieClient represents a Lassie client with its scheme, host, port and LassieClientURL details
type LassieClient struct {
	// Scheme is the protocol scheme (e.g. http, https) used for communication with the Lassie node
	scheme string

	// Host is the hostname or IP address of the Lassie node
	host string

	// Port is the TCP port number used for communication with the Lassie node
	port int

	// LassieClientURL is the URL of the Lassie client
	LassieClientURL string
}

// NewLassieClient creates a new instance of LassieClient with the provided scheme, host, and port.
// It also sets the LassieClientURL based on the provided scheme, host, and port.
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

// GetDataFromCID retrieves the data associated with the given CID from the Lassie node.
//
// The function constructs an HTTP GET request to the LassieClientURL of the LassieClient with the provided CID appended.
// It then sends the request and reads the response body into a byte slice.
//
// Parameters:
//   - cid (string) - The content identifier (CID) of the data to retrieve.
//
// Returns:
//   - data ([]byte) - The retrieved data from the Lassie node.
//   - err (error) - An error if any occurs while constructing the request or reading the response.
func (lassieC *LassieClient) GetDataFromCID(cid string) (data []byte, err error) {
	cidUrl := fmt.Sprintf("%s/%s", lassieC.LassieClientURL, cid)

	req, err := http.NewRequest("GET", cidUrl, nil)
	req.Header.Set("Accept", "*/*")
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: time.Second * 20}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return
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
