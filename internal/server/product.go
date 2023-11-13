package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *EchoServer) GetAllProducts(ctx echo.Context) error {
	vendorID := ctx.QueryParam("vendorId")

	products, err := s.DB.GetAllProducts(ctx.Request().Context(), vendorID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, products)
}
