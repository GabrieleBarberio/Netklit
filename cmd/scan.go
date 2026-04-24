package cmd

import (
	"netklit/pkg/logger"

	"github.com/spf13/cobra"
)

var targetIP string

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Esegue una scansione verso un target",
	Long:  `Avvia una scansione di rete. È necessario specificare un IP di destinazione.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Qui verifichiamo il flag globale e il flag locale
		verbose, _ := cmd.Flags().GetBool("verbose")

		if verbose {
			logger.Log.Info("[DEBUG] Inizializzazione modulo di scansione...")
		}

		if targetIP == "" {
			logger.Log.Info("Errore: devi specificare un target. Usa --help per info.")
			return
		}

		logger.Log.Info("Avviata scansione sul target: %s\n", targetIP)
	},
}

// Un altro blocco statico automatico!
func init() {
	// Aggiungiamo 'scan' come figlio di 'rootCmd'
	rootCmd.AddCommand(scanCmd)

	// Aggiungiamo un flag locale valido SOLO per il comando 'scan'
	// I parametri sono: puntatore alla variabile, nome esteso (--target), nome breve (-t), valore default (""), descrizione
	scanCmd.Flags().StringVarP(&targetIP, "target", "t", "", "IP di destinazione da scansionare")
}
