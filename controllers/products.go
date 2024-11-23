package controllers

import (
	"log"
	"net/http"

	"github.com/Aliz-011/tasukata-go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProducts(c *gin.Context, db *gorm.DB){
	var products []models.Product

	if err := db.Preload("Category").Find(&products).Error; err != nil {
		log.Printf("Error fetching products: %v", err) // Log error for debugging
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"success": true,
		"data": products,
	})
}

func GetProduct(c *gin.Context, db *gorm.DB){
	productId := c.Params.ByName("productId")

	var product models.Product
	if err := db.Where("id = ?", productId).Preload("Category").First(&product).Error; err != nil {
		log.Printf("Error fetching product: %v", err)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found!", 
			"success": false,
		})

    	return
	}

	productResponse := models.ProductResponse {
		ID: product.ID,
		Name: product.Name,
		Description: product.Description,
		Price: int(product.Price),
		Image: product.Image,
		CreatedAt: product.CreatedAt,
		CategoryID: product.CategoryID,
		Category: models.Category{
			ID: product.Category.ID,
			Name: product.Category.Name,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"success": true,
		"data": productResponse,
	})
}