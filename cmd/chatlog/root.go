package chatlog

import (
	"github.com/sjzar/chatlog/internal/chatlog"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	rpcNetwork string
	rpcAddress string
)

func init() {
	// windows only
	cobra.MousetrapHelpText = ""

	rootCmd.PersistentFlags().BoolVar(&Debug, "debug", false, "debug")
	rootCmd.PersistentPreRun = initLog
	rootCmd.Flags().StringVarP(
		&rpcNetwork, "rpc-network", "r", "",
		"The network must be \"tcp\", \"tcp4\", \"tcp6\", \"unix\" or \"unixpacket\".",
	)
	rootCmd.Flags().StringVarP(
		&rpcAddress, "rpc-address", "d", "",
		"exmaple: \":1234\".",
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Err(err).Msg("command execution failed")
	}
}

var rootCmd = &cobra.Command{
	Use:     "chatlog",
	Short:   "chatlog",
	Long:    `chatlog`,
	Example: `chatlog`,
	Args:    cobra.MinimumNArgs(0),
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
	PreRun: initTuiLog,
	Run:    Root,
}

func Root(cmd *cobra.Command, args []string) {

	m, err := chatlog.New("")
	if err != nil {
		log.Err(err).Msg("failed to create chatlog instance")
		return
	}
	m.SetRPC(rpcNetwork, rpcAddress)
	if err := m.Run(); err != nil {
		log.Err(err).Msg("failed to run chatlog instance")
	}

}
