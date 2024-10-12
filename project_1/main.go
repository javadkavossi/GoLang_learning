package main

import (
	"github.com/a-h/templ"
	"github.com/javadkavossi/project_1/views"
	"github.com/labstack/echo/v4"
)

//func Render(c echo.Context, comp templ.Component) error {
//	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
//	return comp.Render(c.Request().Context(), c.Response().Writer)
//}

func Render(c echo.Context, comp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return comp.Render(c.Request().Context(), c.Response().Writer)
}
func main() {
	e := echo.New()
	e.GET("/:name", homeHandler)
	e.Logger.Fatal(e.Start(":3000"))
}

func homeHandler(c echo.Context) error {
	return Render(c, views.Home(c.Param("name")))
}
