package main

import (
	"log"
	"os"

	"github.com/HarisShariff/chainsub/internal/listener"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "chainsub",
		Usage:   "Subscribe to smart contract events on any EVM-compatible chain",
		Version: "0.1.0",
		Commands: []*cli.Command{
			{
				Name:  "listen",
				Usage: "Listen to on-chain events in real time",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "rpc",
						Usage:    "RPC URL of the node",
						Required: true,
						EnvVars:  []string{"RPC_URL"},
					},
					&cli.StringFlag{
						Name:     "contract",
						Usage:    "Contract address",
						Required: true,
						EnvVars:  []string{"CONTRACT_ADDRESS"},
					},
					&cli.StringFlag{
						Name:     "abi",
						Usage:    "Path to ABI JSON file",
						Required: true,
						EnvVars:  []string{"ABI_PATH"},
					},
					&cli.StringFlag{
						Name:     "event",
						Usage:    "Event name (case-sensitive)",
						Required: true,
						EnvVars:  []string{"EVENT_NAME"},
					},
					&cli.StringFlag{
						Name:    "from-block",
						Usage:   "Starting block (default: latest)",
						Value:   "latest",
						EnvVars: []string{"FROM_BLOCK"},
					},
					&cli.StringFlag{
						Name:    "output",
						Usage:   "Output type: stdout, json",
						Value:   "stdout",
						EnvVars: []string{"OUTPUT"},
					},
				},
				Action: func(c *cli.Context) error {
					cfg := listener.Config{
						RPCURL:          c.String("rpc"),
						ContractAddress: c.String("contract"),
						ABIPath:         c.String("abi"),
						EventName:       c.String("event"),
						FromBlock:       c.String("from-block"),
						OutputTarget:    c.String("output"),
					}

					return listener.Listen(cfg)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
