package main

import (
	"log"

	"github.com/hsmtkk/upgraded-invention/wireguardnt"
	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use:  "createadapter adapter-name",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		run(name)
	},
}

func main() {
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}

func run(name string) {
	pool := "WireGuard"
	if err := wireguardnt.CreateAdapter(pool, name); err != nil {
		log.Fatal(err)
	}
}
