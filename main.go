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

	router.Run(":4040")
}

// POST /order
func handleOrder(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

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
