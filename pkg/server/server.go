package server

import (
	"log"
	"net/http"

	"Go-Social/pkg"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Server struct {
	router *mux.Router
	config *root.ServerConfig
}

func NewServer(u root.UserService, p root.PostService, config *root.Config) *Server {
	s := Server{router: mux.NewRouter(), config: config.Server}
	s.router.Use(loggingMiddleware)
	a := authHelper{config.Auth.Secret}
	NewUserRouter(u, s.getSubrouter("/user"), &a)
	NewPostRouter(u, p, s.getSubrouter("/"), &a)
	return &s
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func (s *Server) Start() {
	log.Println("Listening on port " + s.config.Port)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Origin", "X-Requested-With", "Content-Type", "Authorization"},
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
		// Debug:            true,
	}).Handler(s.router)
	if err := http.ListenAndServe(":"+s.config.Port, c); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

func (s *Server) getSubrouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}
