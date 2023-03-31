package main

import (
	"github.com/brilliant-monkey/frigate-kafka-relay/frigate"
	"github.com/brilliant-monkey/frigate-kafka-relay/types"
	"github.com/brilliant-monkey/go-app"
)

const CONFIG_ENVIRONMENT_VARIABLE = "FRIGATE_KAFKA_RELAY_CONFIG_PATH"

func main() {
	app := app.NewApp()
	var c types.Config
	app.LoadConfig(CONFIG_ENVIRONMENT_VARIABLE, &c)

	relay := frigate.NewFrigateRelay(&c)
	app.Go(relay.Start)
	app.Start(func() error {
		return relay.Stop()
	})
}
