package cmd

import (
	"github.com/spf13/cobra"

	"ehkos/pkg/utils"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		broker, _ := cmd.Flags().GetString("broker")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		client_id, _ := cmd.Flags().GetString("client_id")
		topic, _ := cmd.Flags().GetString("topic")

		utils.Subscribe(broker, username, password, client_id, topic)
	},
}

func init() {
	runCmd.AddCommand(serverCmd)

	serverCmd.PersistentFlags().StringP("broker", "b", "", "Possible values:")
	serverCmd.PersistentFlags().StringP("username", "u", "", "Possible values:")
	serverCmd.PersistentFlags().StringP("password", "p", "", "Possible values:")
	serverCmd.PersistentFlags().StringP("client_id", "c", "", "Possible values:")
	serverCmd.PersistentFlags().StringP("topic", "t", "", "Possible values:")

	serverCmd.MarkPersistentFlagRequired("broker")
	serverCmd.MarkPersistentFlagRequired("client_id")
	serverCmd.MarkPersistentFlagRequired("topic")
}
