package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pepeunlimited/go-twirp-starter-kit/internal/server/twirp"
	"github.com/pepeunlimited/go-twirp-starter-kit/pkg/api/v1/services"
	"github.com/pepeunlimited/go-twirp-starter-kit/pkg/env"
	"github.com/pepeunlimited/go-twirp-starter-kit/pkg/middleware"
)

const (
	Version string = "0.0.1"
	Port    int    = 8080
	PortEnv string = "GO_TWIRP_STARTER_KIT_PORT"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Printf("🚀 Started the GoTwirpStarterKit, version=[%s]", Version)
	todoServerPort := env.IntEnv(PortEnv, Port)
	todoServer := services.NewTodoServiceServer(twirp.NewTodoServer(), nil)
	mux := http.NewServeMux()
	mux.Handle(todoServer.PathPrefix(), middleware.Adapt(todoServer))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", todoServerPort), mux); err != nil {
		log.Panic(err)
	}
}
