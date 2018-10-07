package cmd

import (
	"log"
	"os"

	server "github.com/RadicalPixels/server/server"
	"github.com/spf13/cobra"
)

// Start ...
func Start() {
	var (
		nodeURL         string
		contractAddress string
	)
	rootCmd := &cobra.Command{
		Use: "app",
	}
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start the server",
		Long:  "Start the server",
		Run: func(cmd *cobra.Command, args []string) {
			svr := server.NewServer(&server.Config{
				NodeURL:         nodeURL,
				ContractAddress: contractAddress,
			})
			svr.Start()
		},
	}

	startCmd.Flags().StringVarP(&contractAddress, "address", "a", "", "contract address")
	startCmd.Flags().StringVarP(&nodeURL, "node-url", "u", "", "node url")

	rootCmd.AddCommand(startCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
