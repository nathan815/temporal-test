package temporalext

import (
	"log"

	"github.com/nathan815/temporal-hello-world/config"
	"go.temporal.io/sdk/client"
)

func buildClientOptions() client.Options {
	opts := client.Options{
		Namespace: config.TemporalNamespace,
	}

	if len(config.TemporalServerHostPort) > 0 {
		opts.HostPort = config.TemporalServerHostPort
	}

	return opts
}

var ClientOptions = buildClientOptions()

func NewClient() client.Client {
	c, err := client.Dial(ClientOptions)
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	return c
}
