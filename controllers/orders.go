package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Aliz-011/tasukata-go/models"
	"github.com/aidarkhanov/nanoid/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func GetOrders(c *gin.Context, db *gorm.DB){
	var orders []models.Order

	if err := db.Preload("Customer").Preload("OrderItems.Product").Order("created_at desc").Find(&orders).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"success": false,
		})

		return
	}

	var orderResponses []models.OrderResponse
	for _, order := range orders {
		orderResponse := models.OrderResponse{
			ID:      order.ID,
			Address: order.Address,
			Total:   order.Total,
			CustomerID: order.CustomerID,
			Latitude: order.Latitude,
			Longitude: order.Longitude,
			Status: string(order.Status),
			CreatedAt: order.CreatedAt,
			Customer: models.ProfileResponse{
				ID:   order.Customer.ID,
				UserID: order.Customer.UserID,
				Metadata: order.Customer.Metadata,
			},
		}

		for _, item := range order.OrderItems {
			orderResponse.OrderItems = append(orderResponse.OrderItems, models.OrderItemResponse{
				Quantity: item.Quantity,
				ID: item.ID,
				OrderID: item.OrderID,
				ProductID: item.ProductID,
				Product: models.ProductResponse{
					ID: item.Product.ID,
					Name: item.Product.Name,
					Description: item.Product.Description,
					Price: int(item.Product.Price),
					Image: item.Product.Image,
					CreatedAt: item.Product.CreatedAt,
					CategoryID: item.Product.CategoryID,
				},
			})
		}

		orderResponses = append(orderResponses, orderResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"success": true,
		"data": orderResponses,
	})
}

type CreateOrderValues struct {
	CustomerID  string `json:"customerId" binding:"required"`
	Address string `json:"address" binding:"required"`
	Latitude string `json:"latitude" binding:"required"`
	Longitude string `json:"longitude" binding:"required"`
}

func NewOrder(c *gin.Context, db *gorm.DB){
	var values CreateOrderValues
	if err := c.ShouldBindJSON(&values); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "success": false})
		return
	}

	id, err := nanoid.New() //> "i25_rX9zwDdDn7Sg-ZoaH"
	if err != nil {
		log.Fatalln(err)
	}

	lat, _ := strconv.ParseFloat(values.Latitude, 64)
	long, _ := strconv.ParseFloat(values.Longitude, 64)

	order := models.Order{ID: id, CustomerID: values.CustomerID, Latitude: lat, Longitude: long}
	db.Order(order)

	c.JSON(http.StatusOK, gin.H{
		"data": order,
		"message": "Order placed",
		"success": true,
	})
}

func GetOrder(c *gin.Context, db *gorm.DB){
	orderId := c.Params.ByName("orderId")

	var order models.Order
	if err := db.Where("id = ?", orderId).Preload("Customer").Preload("OrderItems.Product").First(&order).Error; err != nil {
		log.Printf("Error fetching order: %v", err)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order not found!", 
			"success": false,
		})

    	return
	}

	orderResponse := models.OrderResponse{
		ID: order.ID,
		Address: order.Address,
		Total: order.Total,
		CustomerID: order.CustomerID,
		Latitude: order.Latitude,
		Longitude: order.Longitude,
		Status: string(order.Status),
		CreatedAt: order.CreatedAt,
		Customer: models.ProfileResponse{
			ID: order.Customer.ID,
			UserID: order.Customer.UserID,
			Metadata: order.Customer.Metadata,
		},
	}

	for _, item := range order.OrderItems {
		orderResponse.OrderItems = append(orderResponse.OrderItems, models.OrderItemResponse{
			Quantity: item.Quantity,
			ID: item.ID,
			OrderID: item.OrderID,
			ProductID: item.ProductID,
			Product: models.ProductResponse{
				ID: item.Product.ID,
				Name: item.Product.Name,
				Description: item.Product.Description,
				Price: int(item.Product.Price),
				Image: item.Product.Image,
				CreatedAt: item.Product.CreatedAt,
				CategoryID: item.Product.CategoryID,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"success": true,
		"data": orderResponse,
	})
}

type UpdateOrderValues struct {
	Status  string `json:"status" binding:"required"`
}

func UpdateOrder(c *gin.Context, db *gorm.DB){
	orderId := c.Params.ByName("orderId")
	
	var order models.Order
	if err := db.Where("id = ?", orderId).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order not found!", 
			"success": false,
		})
		
    	return
	}

	var values UpdateOrderValues
	if err := c.ShouldBindJSON(&values); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	order.Status = models.StatusEnum(values.Status)
	if err := db.Model(&order).Where("id = ?", orderId).Update("status", order.Status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update order",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order updated",
		"success": true,
		"data": order.ID,
	})
}