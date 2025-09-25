package handlers

import (
    "context"
    "net/http"
    "subscriptions-service/internal/models"
    "subscriptions-service/internal/repository"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

type SubscriptionHandler struct {
    repo *repository.SubscriptionRepository
}

func NewSubscriptionHandler(repo *repository.SubscriptionRepository) *SubscriptionHandler {
    return &SubscriptionHandler{repo: repo}
}

// POST /subscriptions
func (h *SubscriptionHandler) Create(c *gin.Context) {
    var req models.Subscription
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    req.ID = uuid.New()

    if err := h.repo.Create(context.Background(), req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, req)
}

// GET /subscriptions/:id
func (h *SubscriptionHandler) GetByID(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    s, err := h.repo.GetByID(context.Background(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }
    c.JSON(http.StatusOK, s)
}

// PUT /subscriptions/:id
func (h *SubscriptionHandler) Update(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    var req models.Subscription
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    req.ID = id

    if err := h.repo.Update(context.Background(), req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, req)
}

// DELETE /subscriptions/:id
func (h *SubscriptionHandler) Delete(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    if err := h.repo.Delete(context.Background(), id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}

// GET /subscriptions/total?user_id=&service_name=
func (h *SubscriptionHandler) GetTotal(c *gin.Context) {
    var userID *uuid.UUID
    if uid := c.Query("user_id"); uid != "" {
        id, err := uuid.Parse(uid)
        if err == nil {
            userID = &id
        }
    }

    var serviceName *string
    if s := c.Query("service_name"); s != "" {
        serviceName = &s
    }

    total, err := h.repo.GetTotal(context.Background(), userID, serviceName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"total": total})
}
