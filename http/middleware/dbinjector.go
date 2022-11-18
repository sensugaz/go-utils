package middleware

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DBInjector(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.WithValue(c.Request().Context(), "db", db.WithContext(c.Request().Context()))
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
