package main

import (
	"github.com/gin-gonic/gin"
	"hsyntnc/shipment-service/order"
	"net/http"
)

// OrderRequest Expected JSON structure.
type OrderRequest struct {
	Quantity int `json:"quantity"`
}

func main() {
	router := gin.Default()
	router.POST("/order", handleOrder)

	router.Run("localhost:8080")
}

// POST /order
func handleOrder(c *gin.Context) {
	var orderReq OrderRequest

	if err := c.BindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalItems, packages := order.FullFillOrder(orderReq.Quantity)

	c.IndentedJSON(http.StatusOK, gin.H{
		"order":      orderReq.Quantity,
		"totalItems": totalItems,
		"packaging":  packages,
	})
}