package handler

import (
	echo "github.com/labstack/echo/v4"
	"net/http"
)

func NewHandler() http.Handler {
	e := echo.New()
    return e
}