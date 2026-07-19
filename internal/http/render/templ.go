package render

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
)

func Templ(c *echo.Context, status int, component templ.Component) error {
	res := c.Response()
	res.Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
	res.WriteHeader(status)
	return component.Render(c.Request().Context(), res)
}
