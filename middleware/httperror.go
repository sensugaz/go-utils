package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sensugaz/go-utils/utils/apperror"
	"net/http"
)

func HTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	if err := apperror.IsError(err); err != nil {
		_ = c.JSON(err.HttpStatus(), err)
		return
	}

	if he, ok := err.(*echo.HTTPError); ok {
		var msg string
		if v, ok := he.Message.(string); ok {
			msg = v
		}

		_ = c.JSON(http.StatusNotFound, apperror.NewError("404", msg))
		return
	}

	_ = c.JSON(http.StatusInternalServerError, apperror.NewError("500", err.Error()))
}
