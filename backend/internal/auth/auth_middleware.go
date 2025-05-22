package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
)

type contextKey string
const userCtxKey contextKey = "user"

func AuthMiddleware(authService domain.AuthService) gin.HandlerFunc {
	return func (c *gin.Context)  {
		// 1. Lê o header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			return
		}

		// 2. Espera o formato "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization Header"})
			return
		}
		tokenString := parts[1]

		// 3. Usa o serviço de autenticação para validar o token
		user, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			return
		}

		// 4. Guarda o usuário autenticado no contexto da requisição
        c.Set(string(userCtxKey), user)

		// 5. Continua o fluxo da requisição
        c.Next()
	}
}