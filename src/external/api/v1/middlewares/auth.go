package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	cognitoJwtVerify "github.com/jhosan7/cognito-jwt-verify"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/external/api/infra/config"
)

func CheckAccessToken(allowedRoutes ...string) gin.HandlerFunc {
	allowed := make(map[string]bool)
	for _, route := range allowedRoutes {
		allowed[route] = true
	}

	return func(c *gin.Context) {
		c.Next()
	}
}

func validateAccessToken(accessToken string) bool {
	cognitoConfig := cognitoJwtVerify.Config{
		UserPoolId: config.GetApiCfg().AuthConfig.UserPoolId,
		ClientId:   config.GetApiCfg().AuthConfig.ClientId,
		TokenUse:   config.GetApiCfg().AuthConfig.TokenUse,
	}

	verify, err := cognitoJwtVerify.Create(cognitoConfig)
	if err != nil {
		fmt.Println(err)
		return false
	}

	_, err = verify.Verify(accessToken)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return false
	}
	return true
}
