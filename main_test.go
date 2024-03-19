package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/order", handleOrder)

	req, _ := http.NewRequest(http.MethodPost, "/order", bytes.NewBufferString(`{"quantity": 6}`))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, 7, int(response["totalItems"].(float64)))

	//expectedPacking := map[string]int{"7": 1}
	//fmt.Println(expectedPacking)
	//fmt.Println(response["packaging"])
	//assert.Equal(t, true, reflect.DeepEqual(expectedPacking, response["packaging"]))
}
