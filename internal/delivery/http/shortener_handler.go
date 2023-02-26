package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	normalizeurl "github.com/sekimura/go-normalize-url"
	"github.com/sotnasr/url-shortener/internal/cache"
	"github.com/sotnasr/url-shortener/internal/utils"
)

const CODE_SIZE = 6

type ShortenerHandler struct {
	cache cache.Cache
}

// NewShortenerHandler function used to configure and register all routees of url shortener.
func NewShortenerHandler(e *echo.Echo, cache cache.Cache) {
	handler := ShortenerHandler{
		cache: cache,
	}

	e.POST("/url-shortener/:url", handler.SetUrl)
	e.GET("/url-shortener/:code", handler.GetUrl)
}

// GetUrl
// @Summary GetUrl
// @Description GetUrl
// @Success 303
// @Router /url-shortener/:code [get]
func (s ShortenerHandler) GetUrl(c echo.Context) error {
	ctx := c.Request().Context()

	code := c.Param("code")
	if code == "" {
		log.Error("code was not found on request")
		return c.NoContent(http.StatusBadRequest)
	}

	url, err := s.cache.Get(ctx, code)
	if err != nil || url == "" {
		log.Error("short url was not found for code: %s", code)
		return c.NoContent(http.StatusNotFound)
	}

	log.Info(fmt.Sprintf("Redirecting client: %s to url: %s", c.RealIP(), url))

	return c.Redirect(http.StatusSeeOther, url)
}

// SetUrl
// @Summary SetUrl
// @Description SetUrl
// @Success 200
// @Router /url-shortener/:url [post]
func (s ShortenerHandler) SetUrl(c echo.Context) error {
	ctx := c.Request().Context()

	url := c.Param("url")
	if url == "" { // TODO: Implements better validation.
		return c.NoContent(http.StatusBadRequest)
	}

	short := utils.RandStringRunes(CODE_SIZE)

	normalized, err := normalizeurl.Normalize(url)
	if err != nil {
		log.Error("error was found while normalize url", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	err = s.cache.Set(ctx, short, normalized, 0)
	if err != nil {
		log.Error("error was found while save short link for url: %s", url)
		return c.NoContent(http.StatusInternalServerError)
	}

	log.Info("short link was created with success for url: %s", url)

	// TODO: Create a model to represents this return.
	return c.JSON(http.StatusOK, map[string]interface{}{"Code": short})
}
