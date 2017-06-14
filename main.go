package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bstick12/onboarding-broker/broker"

	"code.cloudfoundry.org/lager"

	"github.com/gorilla/mux"
	"github.com/pivotal-cf/brokerapi"
)

func main() {
	r := mux.NewRouter()
	broker := broker.Broker{}
	logger := lager.NewLogger("RedisServiceBroker")
	brokerapi.AttachRoutes(r, broker, logger)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
