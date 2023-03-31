package main

import (
	"github.com/brilliant-monkey/frigate-kafka-relay/config"
	"github.com/brilliant-monkey/frigate-kafka-relay/frigate"
	"github.com/brilliant-monkey/go-app"
)

const CONFIG_ENVIRONMENT_VARIABLE = "FRIGATE_KAFKA_RELAY_CONFIG_PATH"

func main() {
	app := app.NewApp()
	var appConfig config.AppConfig
	app.LoadConfig(CONFIG_ENVIRONMENT_VARIABLE, &appConfig)

	relay := frigate.NewFrigateRelay(&appConfig)
	app.Go(relay.Start)
	app.Start(func() error {
		return relay.Stop()
	})
}
