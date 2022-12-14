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

type router func(engine *gin.Engine)

var routers = []router{
	alertRuleRouter,
	alertMetricsRouter,
	alertRecordsRouter,
	alertUsersRouter,
	alertSilenceRouter,
	alertUsersGroupRouter,
}

func NewRouter(c *config.Config) (*Router, error) {
	engine, err := newRouter(c)
	engine.Use(Cors())
	prometheusRouter(engine)
	engine.POST("/login", Login)
	engine.Use(JWTAuth())
	if err != nil {
		return nil, err
	}
	for _, f := range routers {
		f(engine)
	}
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
