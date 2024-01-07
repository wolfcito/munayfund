package restapi

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey string = os.Getenv("SECRETKEY")

func basicTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el token del encabezado de autorización
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autorización faltante"})
			c.Abort()
			return
		}

		// Extraer el token de la cadena "Bearer <token>"
		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no válido"})
			c.Abort()
			return
		}

		// Verificar si el token es válido
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no válido"})
			c.Abort()
			return
		}

		// Si el token es válido, continuar con la solicitud
		c.Next()
	}
}

func protectedHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Esta es una ruta protegida"})
}

func publicHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Esta es una ruta pública"})
}
