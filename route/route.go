package route

import (
	"bus-backend-go/controller"
	"bus-backend-go/keycloak"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Options is a middleware function that appends headers
// for options requests and aborts then exits the middleware
// chain and ends the request.
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}

// Secure is a middleware function that appends security
// and resource access headers.
func Secure(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	//c.Header("X-Frame-Options", "DENY")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}

	// Also consider adding Content-Security-Policy headers
	// c.Header("Content-Security-Policy", "script-src 'self' https://cdnjs.cloudflare.com")
}

type headers struct {
	Cookie        string `json:"Cookie"`
	Authorization string `json:"Authorization"`
}

// 注册路由
func RegisterRoute(engine *gin.Engine, log *logrus.Logger) {
	engine.Use(Options)
	engine.Use(Secure)

	engine.Use(keycloak.Auth(keycloak.LoggedInCheck()))

	AutoRoute(engine, "/api/v1/service", controller.NewServiceController(log))
	AutoRoute(engine, "/api/v1/system", controller.NewSystemController(log))
}