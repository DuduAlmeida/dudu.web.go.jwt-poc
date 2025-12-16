package middlewares

import (
	"net/http"
	"strings"

	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/api/common"
	"github.com/DuduAlmeida/dudu.web.go.jwt-poc/app/services"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(jwtService services.JwtService) echo.MiddlewareFunc {
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

			credentials, err := jwtService.GetAccessCredentialsByToken(accessToken)

			if err != nil || !credentials.Token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token de acesso inválido ou expirado (Manual Validation Failed)")
			}

			ctx.Set(string(common.JwtRegisteredClaimsKey), credentials.RegisteredClaims)

			return next(ctx)
		}
	}
}
