package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/jrpikong/simplebank/db/sqlc"
)

// Server serves HTTP request for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/accounts", server.createAccount)
		v1.GET("accounts/:id", server.getAccount)
		v1.DELETE("/accounts/:id", server.deleteAccount)
		v1.GET("/accounts", server.listAccounts)
		v1.PUT("/accounts", server.updateAccount)
	}

	server.router = router
	return server
}

// Start run the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse set the HTTP error response with header
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
