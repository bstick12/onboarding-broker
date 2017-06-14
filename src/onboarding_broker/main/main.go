package main

import (
	"log"
	"net/http"
	"onboarding_broker/broker"

	"code.cloudfoundry.org/lager"

	"github.com/gorilla/mux"
	"github.com/pivotal-cf/brokerapi"
)

func main() {
	r := mux.NewRouter()
	broker := broker.Broker{}
	logger := lager.NewLogger("RedisServiceBroker")
	brokerapi.AttachRoutes(r, broker, logger)
	log.Fatal(http.ListenAndServe(":8080", r))
}
