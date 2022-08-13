package api

import (
	"github.com/DWHengr/aurora/pkg/config"
	ginlogger "github.com/DWHengr/aurora/pkg/misc/gin"
	"github.com/gin-gonic/gin"
)

type Router struct {
	c      *config.Config
	engine *gin.Engine
}

func NewRouter(c *config.Config) (*Router, error) {
	engine, err := newRouter(c)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	engine.Any("/-/reload")
	return &Router{
		c:      c,
		engine: engine,
	}, nil
}

func newRouter(c *config.Config) (*gin.Engine, error) {

	engine := gin.New()

	engine.Use(ginlogger.LoggerFunc(), ginlogger.RecoveryFunc())

	return engine, nil
}

// Run router.
func (r *Router) Run() {
	r.engine.Run(r.c.Port)
}

// Close router.
func (r *Router) Close() {
}

func (r *Router) router() {
}
