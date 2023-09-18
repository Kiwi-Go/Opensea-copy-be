package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahmanfadhil/gin-bookstore/models"
)

type CreateAccountInput struct {
	NAME     string `json:"name" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
}

type UpdateAccountInput struct {
	NAME     string `json:"name"`
	PASSWORD string `json:"password"`
}

// GET /accounts
// Find all Accounts
func FindAccounts(c *gin.Context) {
	var accounts []models.Account
	models.DB.Find(&accounts)

	c.JSON(http.StatusOK, gin.H{"data": accounts})
}

// GET /accounts/:id
// Find a Account
func FindAccount(c *gin.Context) {
	// Get model if exist
	var account models.Account
	if err := models.DB.Where("id = ?", c.Param("id")).First(&account).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": account})
}

// POST /accounts
// Create new Account
func CreateAccount(c *gin.Context) {
	// Validate input
	var input CreateAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	account := models.Account{NAME: input.NAME, PASSWORD: input.PASSWORD}
	models.DB.Create(&account)

	c.JSON(http.StatusOK, gin.H{"data": account})
}

// PATCH /accounts/:id
// Update a account
func UpdateAccount(c *gin.Context) {
	// Get model if exist
	var account models.Account
	if err := models.DB.Where("id = ?", c.Param("id")).First(&account).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&account).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": account})
}

// DELETE /accounts/:id
// Delete a account
func DeleteAccount(c *gin.Context) {
	// Get model if exist
	var account models.Account
	if err := models.DB.Where("id = ?", c.Param("id")).First(&account).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&account)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
