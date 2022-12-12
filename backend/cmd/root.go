package cmd

import (
	webhookingester "cghendrix/nbfacts/webhooks"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "nbfacts",
	Short: "nbfacts",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			switch args[0] {
			case "server":
				startServer()
			case "worker":
				startPubSubWorker()
			case "local_webhook_handler":
				webhookingester.StartServer()
			}
		}
	},
}

func Execute() {
	viper.AutomaticEnv()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
