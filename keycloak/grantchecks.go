package keycloak

import (
	"bus-backend-go/utils"
	"github.com/gin-gonic/gin"
)

func LoggedInCheck() func(tc *TokenContainer, ctx *gin.Context) bool {
	return func(tc *TokenContainer, ctx *gin.Context) bool {
		username := tc.KeyCloakToken.PreferredUsername
		if len(username) > 0 {
			// 如果顺利拿到了人名就算sso校验通过
			ctx.Set("nowauthorization", ctx.GetHeader("Authorization"))
			ctx.Set("nowuser", tc.KeyCloakToken.PreferredUsername)
			utils.Log.Info(ctx.Get("nowuser"))
			return true
		}
		return false
	}
}
