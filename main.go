package main

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/Aliz-011/tasukata-go/config"
	"github.com/Aliz-011/tasukata-go/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init()  {
	config.LoadEnvVars()
}

func main()  {
	connStr := os.Getenv("DATABASE_URL")
	sqlDB, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to NeonDB") // Log the error with detailed message
		return
	}

	database, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err) // Log the error with detailed message
		return
	}
	
	r := gin.Default()

	r.Use(corsMiddleware())

	r.GET("/api/products", func(ctx *gin.Context) {
		controllers.GetProducts(ctx, database)
	})
	r.GET("/api/products/:productId", func(ctx *gin.Context) {
		controllers.GetProduct(ctx, database)
	})
	
	r.GET("/api/orders", func(ctx *gin.Context) {
		controllers.GetOrders(ctx, database)
	})
	r.POST("/api/orders", func(ctx *gin.Context) {
		controllers.NewOrder(ctx, database)
	})
	r.GET("/api/orders/:orderId", func(ctx *gin.Context) {
		controllers.GetOrder(ctx, database)
	})
	r.PATCH("/api/orders/:orderId", func(ctx *gin.Context) {
		controllers.UpdateOrder(ctx, database)
	})

	r.Run()
}

func corsMiddleware() gin.HandlerFunc {
	// Define allowed origins as a comma-separated string
	originsString := "http://localhost:3000"
	var allowedOrigins []string
	if originsString != "" {
	 // Split the originsString into individual origins and store them in allowedOrigins slice
	 allowedOrigins = strings.Split(originsString, ",")
	}
   
	// Return the actual middleware handler function
	return func(c *gin.Context) {
	 // Function to check if a given origin is allowed
	 isOriginAllowed := func(origin string, allowedOrigins []string) bool {
	  for _, allowedOrigin := range allowedOrigins {
	   if origin == allowedOrigin {
		return true
	   }
	  }
	  return false
	 }
   
	 // Get the Origin header from the request
	 origin := c.Request.Header.Get("Origin")
   
	 // Check if the origin is allowed
	 if isOriginAllowed(origin, allowedOrigins) {
	  // If the origin is allowed, set CORS headers in the response
	  c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	  c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	  c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	  c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	 }
   
	 // Handle preflight OPTIONS requests by aborting with status 204
	 if c.Request.Method == "OPTIONS" {
	  c.AbortWithStatus(204)
	  return
	 }
   
	 // Call the next handler
	 c.Next()
	}
   }