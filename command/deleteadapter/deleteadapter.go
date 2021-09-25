package main

import (
	"log"

	"github.com/hsmtkk/upgraded-invention/wireguardnt"
	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use:  "deleteadapter adapter-name",
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
	handle, err := wireguardnt.OpenAdapter(pool, name)
	if err != nil {
		log.Fatalf("failed to open adapter; %s", err)
	}
	log.Print(handle)
	if err := wireguardnt.DeleteAdapter(handle); err != nil {
		log.Fatalf("failed to delete adapter; %s", err)
	}
}
