package main

import (
    "log"
    "os"
    "subscriptions-service/internal/db"
    "subscriptions-service/internal/handlers"
    "subscriptions-service/internal/repository"

    "github.com/gin-gonic/gin"
)

func main() {
    pool, err := db.ConnectDB()
    if err != nil {
        log.Fatal("DB connection error:", err)
    }
    defer pool.Close()

    repo := repository.NewSubscriptionRepository(pool)
    handler := handlers.NewSubscriptionHandler(repo)

    r := gin.Default()

    r.POST("/subscriptions", handler.Create)
    r.GET("/subscriptions/:id", handler.GetByID)
    r.PUT("/subscriptions/:id", handler.Update)
    r.DELETE("/subscriptions/:id", handler.Delete)
    r.GET("/subscriptions/total", handler.GetTotal)

    port := os.Getenv("APP_PORT")
    r.Run(":" + port)
}
