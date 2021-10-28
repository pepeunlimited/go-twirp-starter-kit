package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pepeunlimited/go-twirp-starter-kit/internal/server/twirp"
	"github.com/pepeunlimited/go-twirp-starter-kit/pkg/api/v1/services"
	"github.com/pepeunlimited/go-twirp-starter-kit/pkg/env"
	"github.com/pepeunlimited/go-twirp-starter-kit/pkg/middleware"

	"go.temporal.io/sdk/client"
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
	temporalClient, err := client.NewClient(client.Options{})
	if err != nil {
		log.Panicf("❌ temporal.client.NewClient occurred issue: %v", err)
	}
	todoServer := services.NewTodoServiceServer(twirp.NewTodoServer(), nil)
	temporalServer := services.NewTemporalServiceServer(twirp.NewTemporalServer(temporalClient), nil)

	mux := http.NewServeMux()
	mux.Handle(todoServer.PathPrefix(), middleware.Adapt(todoServer))
	mux.Handle(temporalServer.PathPrefix(), middleware.Adapt(temporalServer))

	goTwipStarterKitPort := env.IntEnv(PortEnv, Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", goTwipStarterKitPort), mux); err != nil {
		log.Panicf("❌ http.ListenAndServe occurred issue: %v", err)
	}
}
