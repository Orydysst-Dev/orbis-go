package main

import (
	"time"

	"github.com/sourcenetwork/orbis-go/adapter/cobracli"
	hostv1alpha1 "github.com/sourcenetwork/orbis-go/gen/proto/orbis/host/v1alpha1"
	ringv1alpha1 "github.com/sourcenetwork/orbis-go/gen/proto/orbis/ring/v1alpha1"
	secretv1alpha1 "github.com/sourcenetwork/orbis-go/gen/proto/orbis/secret/v1alpha1"
	transportv1alpha1 "github.com/sourcenetwork/orbis-go/gen/proto/orbis/transport/v1alpha1"

	"github.com/NathanBaulch/protoc-gen-cobra/client"
	logging "github.com/ipfs/go-log"
	"github.com/spf13/cobra"
)

var log = logging.Logger("orbis/orbisd")

func main() {

	logging.SetAllLoggers(logging.LevelDPanic)
	logging.SetLogLevelRegex("orbis.*", "info")

	// logging.SetLogLevelRegex("orbis/bulletin.*", "debug")
	// logging.SetLogLevelRegex("orbis/host.*", "debug")
	// loggingv2.SetLogLevel("pubsub", "debug")
	// logging.SetLogLevelRegex("eventbus.*", "debug")
	// golog.SetLogLevelRegex("psrpc.*", "debug")

	err := logging.SetLogLevelRegex("dht/.*", "error")
	if err != nil {
		log.Fatalf("Set log level: %s", err)
	}

	err = logging.SetLogLevelRegex("orbis/transport/.*", "error")
	if err != nil {
		log.Fatalf("Set log level: %s", err)
	}

	rootCmd := &cobra.Command{
		Use:          "orbisd",
		Long:         "Orbis is a hybrid secrets management engine designed as a decentralized custodial system.",
		SilenceUsage: true,
	}

	// Setup the start command for the Orbis server.
	startCmd, err := cobracli.StartCmd(setupServer)
	if err != nil {
		log.Fatalf("Start command: %s", err)
	}

	// Setup client commands for the Orbis client.
	rootCmd.AddCommand(startCmd)

	opts := []client.Option{
		client.WithTimeout(1 * time.Second),
	}
	rootCmd.AddCommand(
		hostv1alpha1.HostServiceClientCommand(opts...),
		ringv1alpha1.RingServiceClientCommand(opts...),
		secretv1alpha1.SecretServiceClientCommand(opts...),
		transportv1alpha1.TransportServiceClientCommand(opts...),
	)

	rootCmd.Execute() // nolint:errcheck
}
