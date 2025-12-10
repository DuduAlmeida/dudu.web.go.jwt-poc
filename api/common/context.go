package common

import "github.com/labstack/echo/v4"

const (
	UserClaimsKey = "userClaims"
)

func GetEchoContextValue[T any](ctx echo.Context, key string) *T {
	pointer := ctx.Get(string(key))

	if pointer == nil {
		return nil
	}

	value, ok := pointer.(*T)
	if !ok {
		return nil
	}

	return value
}
