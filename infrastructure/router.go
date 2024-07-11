package infrastructure

import (
	"golang_test/adapter/controller"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/invoice/customer", func(ctx echo.Context) error {
		resp, err := c.Invoice.GetCustomerList(ctx)
		if err != nil {
			report, ok := err.(*echo.HTTPError)
			if !ok {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			ctx.Logger().Error(report)
			return ctx.JSON(report.Code, report)
		}
		return ctx.JSON(http.StatusOK, resp)
	})

	e.GET("/invoice/item", func(ctx echo.Context) error {
		resp, err := c.Invoice.GetItemList(ctx)
		if err != nil {
			report, ok := err.(*echo.HTTPError)
			if !ok {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			ctx.Logger().Error(report)
			return ctx.JSON(report.Code, report)
		}
		return ctx.JSON(http.StatusOK, resp)
	})

	e.GET("/invoice", func(ctx echo.Context) error {
		resp, err := c.Invoice.GetInvoiceList(ctx)
		if err != nil {
			report, ok := err.(*echo.HTTPError)
			if !ok {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			ctx.Logger().Error(report)
			return ctx.JSON(report.Code, report)
		}
		return ctx.JSON(http.StatusOK, resp)
	})

	e.POST("/invoice", func(ctx echo.Context) error {
		resp, err := c.Invoice.SaveInvoiceTrx(ctx)
		if err != nil {
			report, ok := err.(*echo.HTTPError)
			if !ok {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			ctx.Logger().Error(report)
			return ctx.JSON(report.Code, report)
		}
		return ctx.JSON(http.StatusOK, resp)
	})

	e.GET("/invoice/detail", func(ctx echo.Context) error {
		resp, err := c.Invoice.GetInvoiceById(ctx)
		if err != nil {
			report, ok := err.(*echo.HTTPError)
			if !ok {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			ctx.Logger().Error(report)
			return ctx.JSON(report.Code, report)
		}
		return ctx.JSON(http.StatusOK, resp)
	})

	return e
}
