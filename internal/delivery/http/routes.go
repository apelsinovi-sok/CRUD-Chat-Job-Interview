package http

import "CRUD-Chat-Test-Task/internal/delivery/http/internal"

func (s *httpServer) setRoutes() {
	s.setApi()
}

func (s *httpServer) setApi() {
	s.router.POST("/create", internal.CreateChat)
	s.router.POST("/add-message", internal.AddMessage)
	s.router.GET("/chats/:name/:limit", internal.GetListMessages)
	s.router.GET("/chat/:id", internal.GetMessage)
}
