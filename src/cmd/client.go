package cmd

import (
	"ehkos/pkg/utils"

	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		broker, _ := cmd.Flags().GetString("broker")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		client_id, _ := cmd.Flags().GetString("client_id")
		topic, _ := cmd.Flags().GetString("topic")

		utils.Publish(file, broker, username, password, client_id, topic)
	},
}

func init() {
	runCmd.AddCommand(clientCmd)

	clientCmd.PersistentFlags().StringP("file", "f", "", "Possible values:")
	clientCmd.PersistentFlags().StringP("broker", "b", "", "Possible values:")
	clientCmd.PersistentFlags().StringP("username", "u", "", "Possible values:")
	clientCmd.PersistentFlags().StringP("password", "p", "", "Possible values:")
	clientCmd.PersistentFlags().StringP("client_id", "c", "", "Possible values:")
	clientCmd.PersistentFlags().StringP("topic", "t", "", "Possible values:")

	clientCmd.MarkPersistentFlagRequired("file")
	clientCmd.MarkPersistentFlagRequired("broker")
	clientCmd.MarkPersistentFlagRequired("client_id")
	clientCmd.MarkPersistentFlagRequired("topic")
}
