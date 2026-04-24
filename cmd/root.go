package cmd

import (
	"netklit/pkg/logger"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "Netklit",
	Short: "Netklit - CLI per operazioni di rete",
	Long:  `Applicazione da riga di comando per gestire i processi di rete, ottimizzarli e rilanciarli`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Log.Info("Benvenuto in Netkit! Usa --help per vedere cosa posso fare.")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logger.Log.Error(err)
		os.Exit(1)
	}
}
func init() {
	// Qui definire dei flag GLOBALI (es. --verbose) che valgono per tutti i comandi
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "Attiva l'output prolisso")
}
