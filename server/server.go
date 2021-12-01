package server

import (
"context"
	"log"
	"net/http"

"github.com/labstack/echo/v4"
"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	VersionInfo
	fs *http.Handler
	port string
}

type VersionInfo struct {
	Version string
	Commit  string
	Build   string
}

func New(info VersionInfo, port string, fs *http.Handler) *Server {
	return &Server{
		VersionInfo: info,
		fs: fs,
		port:        port,
	}
}


func (s Server) Serve(ctx context.Context) error {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(middleware.Recover())
	s.initHandlers(e)

	go func() {
		e.Logger.Infof("start server on port: %s", s.port)
		err := e.Start(":" + s.port)
		if err != nil {
			e.Logger.Errorf("start server error: %v", err)
		}
	}()

	<-ctx.Done()

	return e.Shutdown(ctx)
}

func (s Server) initHandlers(e *echo.Echo) {

	e.GET("/", echo.WrapHandler(*s.fs))
	e.GET("/__heartbeat__", heartbeatHandler)
	e.GET("/__version__", s.versionHandler)
	e.Any("/*", func(c echo.Context) error {
		return c.NoContent(http.StatusNotFound)
	})
}

func handler(c echo.Context) error {
	log.Printf("Hello %v", c.Path())
	return c.String(http.StatusOK, "Hello, World! Welcome to GeekBrains!\n")
}


