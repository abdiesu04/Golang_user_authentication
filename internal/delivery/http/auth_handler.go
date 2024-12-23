package http

import (
	"net/http"
	"user/internal/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, authUsecase *usecase.AuthUsecase) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", func(c *gin.Context) {
			var req struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
				return
			}

			err := authUsecase.Register(req.Username, req.Password)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, gin.H{"message": "user registered"})
		})

		authGroup.POST("/login", func(c *gin.Context) {
			var req struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
				return
			}

			tokens, err := authUsecase.Login(req.Username, req.Password)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, tokens)
		})
	}
}
