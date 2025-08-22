package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc{
	return func(c *gin.Context){
		//pega os claims que foram salvos no contexto 
		claimsInterface, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "claims não encontrados"})
			c.Abort()
			return
		}

		//Conversão
		claims, ok := claimsInterface.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato invalido dos claims"})
			c.Abort()
			return
		}

		//pegar role do token
		role, ok := claims["role"].(string)
		if !ok{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role não encontrada"})
			c.Abort()
			return
		}

		//verificar se a roles está entre as roles permitidas
		for _, allowed := range allowedRoles{
			if role == allowed{
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Permissão negada"})
		c.Abort()	
	}
}