package server

import (
	"context"
	"fmt"
	"github.com/Gregory-Development/http-router/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	HttpServer *http.Server
	Config *config.Config
}

func NewServer(config *config.Config) *Server {
	var S Server

	S.Config = config
	S.HttpServer = S.newHttpServer()

	return &S
}

func (S *Server) newHttpServer() *http.Server {
	cfg := S.Config
	svr := &http.Server{
		Addr: fmt.Sprintf("%s:%d", cfg.HttpIPv4BindAddress, cfg.HttpIPv4BindPort),
		Handler: S.newRouter(),
		ReadTimeout: cfg.HttpReadTimeout * time.Second,
		WriteTimeout: cfg.HttpWriteTimeout * time.Second,
		IdleTimeout: cfg.HttpIdleTimeout * time.Second,
	}
	return svr
}

func (S *Server) newRouter() *mux.Router {
	rtr := mux.NewRouter()
	return rtr
}

func (S *Server) Run() {
	cfg := S.Config

	log.Printf("starting http server bound to %s:%d\n", cfg.HttpIPv4BindAddress, cfg.HttpIPv4BindPort)

	go func(){
		if err := S.HttpServer.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<- c

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("shutting down the http server")
	err := S.HttpServer.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}
}