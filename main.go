package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
)

type item struct {
	ShortDescription string `json: "shortDescription"`
	Price            string `json: "price"`
}

type receipt struct {
	ID           string `json: "id"`
	Retailer     string `json: "retailer"`
	PurchaseDate string `json: "purchaseDate"`
	PurchaseTime string `json: "purchaseTime"`
	Items        []item `json: "items`
	Total        string `json: "total"`
}

var receipts = []receipt{}

// RULES
func isAlphanumeric(s string) int {
	result := ""
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			result += string(char)
		}
	}
	fmt.Printf("result: %v\n", result)
	return len(result)
}

func totalIsRoundDollar(total string) int16 {
	num, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return 0
	}

	if num == float64(int(num)) {
		return 50
	} else {
		return 0
	}
}

func totalIsMultiple(total string) int {
	num, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return 0
	}

	if math.Mod(num, 0.25) == 0 {
		return 25
	} else {
		return 0
	}
}

func receiptItems(items []item) float64 {
	pairs := math.Floor(float64(len(items)) / 2)
	result := pairs * 5
	return result
}

func itemDescriptionIsMultiple(items []item) float64 {
	var points = 0.00
	for i := 0; i < len(items); i++ {
		if len(strings.TrimSpace(items[i].ShortDescription))%3 == 0 {
			num, err := strconv.ParseFloat(items[i].Price, 64)

			if err != nil {
				return 0
			}

			points = math.Round(num * 0.2)
		}
	}
	return points
}

func purchaseDayOdd(purchaseDate string) int {
	d, err := time.Parse("15:04", purchaseDate)
	if err != nil {
		return 0
	}

	if d.Day()%2 == 0 {
		return 6
	} else {
		return 0
	}
}

func purchaseTimeWithinRange(purchaseTime string) int {
	inputTime, err := time.Parse("15:04", purchaseTime)
	if err != nil {
		return 0
	}

	start, _ := time.Parse("15:04", "14:00")
	end, _ := time.Parse("15:04", "16:00")

	if inputTime.After(start) && inputTime.Before(end) {
		return 10
	} else {
		return 0
	}
}

// HANDLERS
func getReceipts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, receipts)
}

func processReceipt(c *gin.Context) {
	var newReceipt receipt

	if err := c.BindJSON(&newReceipt); err != nil {
		return
	}

	newReceipt.ID = strconv.FormatInt(time.Now().UnixNano(), 10)

	receipts = append(receipts, newReceipt)
	c.IndentedJSON(http.StatusCreated, gin.H{"id": newReceipt.ID})
}

func getPoints(c *gin.Context) {
	id := c.Param("id")

	// Loop through to find receipt that maitche whose ID matches id
	for _, receipt := range receipts {
		if receipt.ID == id {
			var points = isAlphanumeric(receipt.Retailer) + int(totalIsRoundDollar(receipt.Total)) + totalIsMultiple(receipt.Total) + int(receiptItems(receipt.Items)) + int(itemDescriptionIsMultiple(receipt.Items)) + purchaseDayOdd(receipt.PurchaseDate) + purchaseTimeWithinRange(receipt.PurchaseTime)
			c.IndentedJSON(http.StatusOK, gin.H{"points": points})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "receipt not found"})
}

func main() {
	router := gin.Default()
	router.GET("/receipts", getReceipts)
	router.POST("/receipts/process", processReceipt)
	router.GET("/receipts/:id/points", getPoints)
	router.Run("localhost:8080")
}
