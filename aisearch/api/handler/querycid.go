package handler

import (
	"log"
	"net/http"

	"ssaisearch/lassie" // Assuming the package name is "lassie"

	"github.com/gin-gonic/gin"
)

// QueryCIDHandler handles the GET request to retrieve and return the JSON content of a CID.
func QueryCIDHandler(c *gin.Context) {
	// Extract the CID from the path parameter.
	cid := c.Param("cid")
	if cid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CID parameter is required"})
		return
	}

	// Create a new IPFS client.
	ipfsC := lassie.NewLassieClient("http",
		"127.0.0.1",
		62156)

	// Retrieve the data from the CID.
	D, err := lassie.GetJSONDATAFromCID(ipfsC, cid)
	if err != nil {
		log.Panicln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data from CID"})
		return
	}

	// Assuming the first file in D is the JSON content we want to return.
	if len(D) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No files found in the CID"})
		return
	}

	// Return the raw JSON data.
	c.Data(http.StatusOK, "application/json", []byte(D))
}
