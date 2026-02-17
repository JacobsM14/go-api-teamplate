package api

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go-api-template/api/middlewares"
	db "go-api-template/database"
)

type Server struct {
	listenAddress string
	store         db.Storage
}

func NewServer(listenAddress string, store db.Storage) *Server {
	return &Server{
		listenAddress: listenAddress,
		store:         store,
	}
}

func (s *Server) Start() error {
	router := http.NewServeMux()

	// Set routes
	s.setRoutes(router)

	server := &http.Server{
		Addr:    s.listenAddress,
		Handler: middlewares.CORS(router),
	}

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Shutting down server...")
		server.Close()
	}()

	log.Printf("Server running on port %s", s.listenAddress)
	return server.ListenAndServe()
}

func (s *Server) setRoutes(router *http.ServeMux) {
	// authMiddleware := middlewares.JWTAuthMiddleware([]byte(os.Getenv("SECRET")), s.store)

}
