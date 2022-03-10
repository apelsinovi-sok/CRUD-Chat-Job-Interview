package http

import (
	"CRUD-Chat-Test-Task/internal/core/interfaces"
	"fmt"
	"github.com/gin-gonic/gin"
)

type httpServer struct {
	router  *gin.Engine
	usecase interfaces.Usecase
}

func New(usecase interfaces.Usecase) *httpServer {
	s := &httpServer{
		router:  gin.Default(),
		usecase: usecase,
	}
	s.setExternalParamInHandlers()
	s.setRoutes()
	return s
}

func (s *httpServer) setExternalParamInHandlers() {
	s.router.Use(func(c *gin.Context) {
		c.Set("usecase", &s.usecase)
		c.Next()
	})
}

func (s *httpServer) Run(port string) {
	port = fmt.Sprintf(":%s", port)
	err := s.router.Run(port)
	if err != nil {
		panic(err.Error())
	}
}
