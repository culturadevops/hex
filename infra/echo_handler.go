package infra

import (
	"github.com/labstack/echo"
)

type EchoHandler interface {
	NewHandler(s Service) EchoHandler {
	ListAllFor(c echo.Context) error
	Get(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
}
