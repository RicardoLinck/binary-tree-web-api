package network

import (
	"net/http"
	"sync"

	"github.com/RicardoLinck/binary-tree-web-api/binarytree"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type srv struct {
	logger *zap.Logger
	bt     binarytree.BinaryTree
	m      sync.Mutex
}

// BuildServer builds and returns a new http server
func BuildServer(l *zap.Logger) *http.Server {
	return &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: buildHandler(l),
	}
}

func buildHandler(l *zap.Logger) *mux.Router {
	r := mux.NewRouter()
	s := srv{logger: l}
	r.Use(correlationIDMiddleware)
	r.Use(s.loggingMiddleware)
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/bt/insert/{val}", s.insertHandler).Methods("POST")
	r.HandleFunc("/bt/contains/{val}", s.containsHandler).Methods("GET")
	r.HandleFunc("/bt", s.getHandler).Methods("GET")
	return r
}
