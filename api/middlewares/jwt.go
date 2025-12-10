package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/api/common"
	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/domain/user"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(accessTokenSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			authHeader := ctx.Request().Header.Get("Authorization")

			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token de acesso não fornecido")
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Formato do token inválido (esperado: Bearer <token>)")
			}
			accessToken := parts[1]

			claims := new(user.UserClaims)
			token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
				}
				return []byte(accessTokenSecret), nil
			})

			if err != nil || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token de acesso inválido ou expirado (Manual Validation Failed)")
			}

			ctx.Set(string(common.UserClaimsKey), claims)

			return next(ctx)
		}
	}
}
