package middleware

import (
	utils "github.com/TTaket/gochat/utils"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware
// JWT auth middleware
// jwt in authorization header
// check token in redis by username
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get jwt from header
		jwtTokenString := c.GetHeader("Authorization")
		if jwtTokenString == "" {
			c.JSON(401, gin.H{
				"message": "no jwt authorization in header",
			})
			c.Abort()
			return
		}

		// verify jwt
		ok, err := utils.VerifyToken(jwtTokenString)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "server error",
				"error":   err.Error(),
				"getjwt":  jwtTokenString,
			})
			c.Abort()
			return
		}
		if !ok {
			c.JSON(401, gin.H{
				"message": "jwtTokenString is not valid",
			})
			c.Abort()
			return
		}

		// parse jwt
		jwtInfo, err := utils.ParseToken(jwtTokenString)
		if err != nil {
			c.JSON(500, gin.H{
				"message":    "server error",
				"error":      err.Error(),
				"parse info": jwtInfo,
			})
			c.Abort()
			return
		}
		if jwtInfo == nil || jwtInfo.Username == "" {
			c.JSON(401, gin.H{
				"message": "parse info is empty or no username in parse info",
			})
			c.Abort()
			return
		}

		// check redis
		jwtInRedis, err := utils.GetRedisClient().Get(jwtInfo.Username).Result()
		if err != nil {
			c.JSON(500, gin.H{
				"message": "server error",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}
		if jwtInRedis != jwtTokenString {
			c.JSON(401, gin.H{
				"message": "jwt is not equal to redis",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
