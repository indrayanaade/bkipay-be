package main

import (
	"log"
	
	"github.com/gin-gonic/gin"
	"github.com/indrayanaade/bkipay-api/internal/database"
	"github.com/indrayanaade/bkipay-api/internal/redis"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	//Init Redis client di awal
	if err := redis.InitRedis(); err != nil {
		log.Fatalf("Redis init failed: %v", err)
	}

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "BKIPay API Root"})
	})

	r.GET("/api/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from BKIPay API!"})
	})


	// Cek koneksi DB
	r.GET("/api/ping-db", func(c *gin.Context) {
		db, err := database.ConnectDB()
		if err != nil {
			c.JSON(500, gin.H{"status": "error", "message": err.Error()})
			return
		}
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(500, gin.H{"status": "error", "message": err.Error()})
			return
		}
		if err := sqlDB.Ping(); err != nil {
			c.JSON(500, gin.H{"status": "error", "message": "DB not reachable"})
			return
		}
		c.JSON(200, gin.H{"status": "ok", "message": "DB connected successfully"})
	})

	// Cek koneksi Redis
	r.GET("/api/ping-redis", func(c *gin.Context) {
		ctx := c.Request.Context()
		if redis.RedisClient == nil {
			c.JSON(500, gin.H{"status": "error", "message": "Redis client not initialized"})
			return
		}
		if err := redis.RedisClient.Ping(ctx).Err(); err != nil {
			c.JSON(500, gin.H{"status": "error", "message": "Redis not reachable", "error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "ok", "message": "Redis connected successfully"})
	})	

	r.Run("0.0.0.0:8010")
}