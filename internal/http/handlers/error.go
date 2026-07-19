package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/trigex/trigex.moe/internal/views"
)

func NewErrorHandler(siteName string) echo.HTTPErrorHandler {
	return func(c *echo.Context, err error) {
		if resp, uErr := echo.UnwrapResponse(c.Response()); uErr == nil {
			if resp.Committed {
				return // already sent by a handler / middleware
			}
		}

		code := http.StatusInternalServerError
		var sc echo.HTTPStatusCoder
		if errors.As(err, &sc) {
			if tmp := sc.StatusCode(); tmp != 0 {
				code = tmp
			}
		}

		var cErr error
		if c.Request().Method == http.MethodHead {
			cErr = c.NoContent(code)
		} else {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
			c.Response().WriteHeader(code)

			cErr = views.Layout(
				fmt.Sprintf("%s | %d Error", siteName, code),
				views.ErrorPage(code, http.StatusText(code)),
			).Render(c.Request().Context(), c.Response())
		}
		if cErr != nil {
			c.Logger().Error("failed to send error page", "error", errors.Join(err, cErr))
		}
	}
}
