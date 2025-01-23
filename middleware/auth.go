package middleware

import (
	utils "github.com/TTaket/gochat/utils"
	"github.com/gin-gonic/gin"
	viper "github.com/spf13/viper"
)

// AuthMiddleware
// JWT auth middleware
// jwt in authorization header
// check token in redis by username
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get jwt from header
		jwtTokenString := c.GetHeader("Authorization")
		if jwtTokenString == "" || len(jwtTokenString) < len("Bearer ") {
			c.JSON(401, gin.H{
				"message": "no jwt authorization in header",
			})
			c.Abort()
			return
		}
		jwtTokenString = jwtTokenString[len("Bearer "):]

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

// AuthMiddleware
// Admin auth middleware
// check admin in hard code
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get jwt from header
		jwtTokenString := c.GetHeader("Authorization")
		if jwtTokenString == "" || len(jwtTokenString) < len("Bearer ") {
			c.JSON(401, gin.H{
				"message": "no jwt authorization in header",
			})
			c.Abort()
			return
		}
		jwtTokenString = jwtTokenString[len("Bearer "):]

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

		// check admin
		adminMapTemp := viper.GetStringMap("admin")
		var adminMap = make(map[string]bool)
		for _, v := range adminMapTemp {
			if str, ok := v.(string); ok {
				adminMap[str] = true
			}
		}
		if _, ok := adminMap[jwtInfo.Username]; !ok {
			c.JSON(401, gin.H{
				"message":     "you are user but not admin",
				"yoursername": jwtInfo.Username,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
