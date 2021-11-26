package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func heartbeatHandler(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (s Server) versionHandler(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		map[string]string{
			"version": s.VersionInfo.Version,
			"commit":  s.VersionInfo.Commit,
			"build":   s.VersionInfo.Build,
		},
	)
}


