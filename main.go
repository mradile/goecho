package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("*", func(c echo.Context) error {
		sb := strings.Builder{}
		sb.WriteString(fmt.Sprintf("Path: %s\n", c.Request().URL.String()))
		sb.WriteString(fmt.Sprintf("Host: %s\n", c.Request().Host))

		sb.WriteString("\n")

		sb.WriteString("Header:\n")
		h := c.Request().Header
		for k, v := range h {
			sb.WriteString(fmt.Sprintf("  %s: ", k))
			sb.WriteString(fmt.Sprintf("%s", strings.Join(v, ",")))
			sb.WriteString("\n")
		}

		sb.WriteString("\n")

		return c.String(http.StatusOK, sb.String())
	})

	e.Logger.Fatal(e.Start(":3000"))
}
