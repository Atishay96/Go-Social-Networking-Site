package server

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	"Go-Social/pkg"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	config *root.ServerConfig
}

func NewServer(u root.UserService, p root.PostService, config *root.Config) *Server {
	s := Server{router: mux.NewRouter(), config: config.Server}

	a := authHelper{config.Auth.Secret}
	NewUserRouter(u, s.getSubrouter("/user"), &a)
	NewPostRouter(u, p, s.getSubrouter("/"), &a)
	return &s
}

func (s *Server) Start() {
	log.Println("Listening on port " + s.config.Port)
	headers := handlers.AllowedHeaders([]string{"*"})
	methods := handlers.AllowedMethods([]string{"*"})
	origins := handlers.AllowedOrigins([]string{"*"})
	// handlers := simpleChain(handlers.LoggingHandler(os.Stdout, s.router), handlers.CORS(handlers.AllowedHeaders([]string{"Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(s.router))
	if err := http.ListenAndServe(":"+s.config.Port, handlers.CORS(headers, methods, origins)(s.router)); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

func (s *Server) getSubrouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}
