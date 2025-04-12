package utils

// Package utils provides utility functions for the AI search system.
// This file contains Gin web framework related utilities.

import "github.com/gin-gonic/gin"

// GinErrorMsg sends a JSON error response with status code 500.
//
// Parameters:
//   c: Gin context
//   err: Error to send in response
//
// Example:
//   if err != nil {
//       utils.GinErrorMsg(c, err)
//       return
//   }
func GinErrorMsg(c *gin.Context, err error) {
	c.JSON(500, gin.H{
		"message": err.Error(),
	})
}

// GinErrorMsgTxt sends a JSON error response with status code 500 using a string message.
//
// Parameters:
//   c: Gin context
//   err: Error message to send
//
// Example:
//   utils.GinErrorMsgTxt(c, "Invalid request parameters")
func GinErrorMsgTxt(c *gin.Context, err string) {
	c.JSON(500, gin.H{
		"message": err,
	})
}
