package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sotnasr/url-shortener/internal/cache"
)

type ShortenerHandler struct {
	cache cache.Cache
}

// GetUrl
// @Summary GetUrl
// @Description GetUrl
// @Success 303
// @Router /url-shortener/:code [get]
func (s *ShortenerHandler) GetUrl(c echo.Context) error {
	ctx := c.Request().Context()

	code := c.Param(":code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}

	url, err := s.cache.Get(ctx, code)
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	c.Redirect(http.StatusSeeOther, url)
	return nil
}
