package main

import (
	"git.brilliantmonkey.net/frigate/frigate-clips/app"
	"git.brilliantmonkey.net/frigate/frigate-clips/frigate"
	"git.brilliantmonkey.net/frigate/frigate-clips/types"
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
