package handler

import (
	"log"
	"net/http"

	"ssaisearch/lassie" // Assuming the package name is "lassie"

	"github.com/gin-gonic/gin"
)

// QueryCIDHandler handles CID-based content retrieval from IPFS network
// Workflow:
// 1. Extracts CID from request path
// 2. Initializes IPFS client connection
// 3. Fetches JSON data from IPFS using the CID
// 4. Returns raw JSON response or appropriate error
func QueryCIDHandler(c *gin.Context) {
	// Get CID parameter from URL path
	cid := c.Param("cid")
	if cid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CID parameter is required"})
		return
	}

	// Initialize IPFS client with default local node configuration
	ipfsC := lassie.NewLassieClient("http",
		"127.0.0.1",
		62156)

	// Fetch JSON data from IPFS network using the CID
	jsonData, err := lassie.GetJSONDATAFromCID(ipfsC, cid)
	if err != nil {
		log.Panicln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data from CID"})
		return
	}

	// Verify we received valid JSON data
	if len(jsonData) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No files found in the CID"})
		return
	}

	// Stream the raw JSON response back to client
	c.Data(http.StatusOK, "application/json", []byte(jsonData))
}
