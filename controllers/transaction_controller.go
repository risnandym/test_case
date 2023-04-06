package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"test-case/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllTransactions godoc
// @Summary Get all Transaction.
// @Description Get a list of Transaction.
// @Tags Transaction
// @Produce json
// @Success 200 {object} []models.Transaction
// @Router /transactions [get]
func GetAllTransaction(c *gin.Context) {

	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var merks []models.Transaction
	db.Find(&merks)

	c.JSON(http.StatusOK, gin.H{"data": merks})
}

// CreateTransactions godoc
// @Summary Create New Transaction.
// @Description Creating a new Transaction.
// @Tags Transaction
// @Param Body body models.TransactionRequest true "the body to create a new Transaction"
// @Produce json
// @Success 200 {object} models.Transaction
// @Router /save-transactions [post]
func SaveTransaction(c *gin.Context) {
	start := time.Now()
	db := c.MustGet("db").(*gorm.DB)
	var request models.TransactionRequest
	// var trans models.Transaction

	// Validate input
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactions := request.Data
	for idx := range transactions {
		transactions[idx].RequestID = request.RequestID
	}

	// insert data using goroutine
	var wg sync.WaitGroup
	wg.Add(10)
	go func(values []models.Transaction) {
		defer wg.Done()

		db.CreateInBatches(&values, 100)

	}(transactions)

	elapsed := time.Since(start).Seconds()

	c.JSON(http.StatusOK, gin.H{"time_elapsed": elapsed, "succes": fmt.Sprintf("save %d transaction", len(request.Data))})
}

// GetDummyTransactions godoc
// @Summary Create Dummy Transaction.
// @Description Creating a new dummy Transaction.
// @Tags Transaction
// @Param dummy path string true "Total of Dummy Transaction required"
// @Produce json
// @Success 200 {object} models.TransactionRequest
// @Router /dummy-transactions/{dummy} [get]
func GenerateDummy(c *gin.Context) {
	var data []models.Transaction
	var response models.TransactionRequest

	total, _ := strconv.Atoi(c.Param("dummy"))

	response.RequestID = rand.Intn(100)

	for i := 1; i <= total; i++ {
		var trans models.Transaction
		trans.Customer = fmt.Sprintf("Cusomer-%d", i)
		trans.Quantity = rand.Intn(100)
		trans.Price = rand.Float64()
		trans.TimeStamp = time.Now()
		data = append(data, trans)
	}

	response.Data = data
	// Save Transaction

	c.JSON(http.StatusOK, response)
}
