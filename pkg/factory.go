package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/kailash-bhanushali/backend-golang/internal/config"
)

type Factory interface {
	CreateRouterEngine() *gin.Engine
	EnvRelation() string
}

type FactoryImpl struct {
	routerEngine *gin.Engine
	env          string
}

func NewFactory(config *config.ServerConfig) Factory {
	return &FactoryImpl{
		routerEngine: gin.Default(),
		env:          config.Env,
	}
}

func (f *FactoryImpl) CreateRouterEngine() *gin.Engine {
	return f.routerEngine
}

func (f *FactoryImpl) EnvRelation() string {
	return f.env
}
