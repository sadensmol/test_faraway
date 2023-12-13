package main

import (
	client "sadensmol/go/test_faraway/internal/client"
	server "sadensmol/go/test_faraway/internal/server"

	"github.com/spf13/cobra"
)

var (
	port int
	url  string

	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start the server",
		Run: func(cmd *cobra.Command, args []string) {
			server.NewServer(port).Start()
		},
	}

	clientCmd = &cobra.Command{
		Use:   "client",
		Short: "Start a client and send a request to the server",
		Run: func(cmd *cobra.Command, args []string) {
			client.NewClient(url).Start()
		},
	}
)

func main() {
	var rootCmd = &cobra.Command{Use: "test_faraway"}
	serverCmd.Flags().IntVarP(&port, "port", "p", 8080, "port to listen on")
	clientCmd.Flags().StringVarP(&url, "url", "u", "http://localhost:8080", "url of the server")
	rootCmd.AddCommand(serverCmd, clientCmd)
	rootCmd.Execute()
}
