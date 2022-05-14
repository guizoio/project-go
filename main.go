package main

import (
	"awesomeProject1/src/cmd"
	"awesomeProject1/src/infrastructure"
	"context"
	"github.com/spf13/cobra"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer stop()

	config := infrastructure.NewConfig()
	containerDI := infrastructure.NewContainerDI(config)
	defer containerDI.ShutDown()

	cmdHttpServer := &cobra.Command{
		Use:   "httpserver",
		Short: "Run httpserver",
		Run: func(cli *cobra.Command, args []string) {
			cmd.StartHttp(ctx, containerDI)
		},
	}

	var rootCmd = &cobra.Command{Use: "APP"}
	rootCmd.AddCommand(cmdHttpServer)
	rootCmd.Execute()
}
