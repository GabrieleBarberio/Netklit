package scan

import (
	"fmt"
	"net"
	"netklit/cmd"
	"netklit/internal/config"
	"netklit/pkg/logger"
	"time"

	"github.com/spf13/cobra"
)

var targetIP string
var targetPort int

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Esegue una scansione verso un target",
	Long:  `Avvia una scansione di rete. È necessario specificare un IP di destinazione.`,
	Run: func(cobraCmd *cobra.Command, args []string) {
		verbose, _ := cobraCmd.Root().PersistentFlags().GetBool("verbose")
		if verbose {
			logger.Log.Info("[DEBUG] Inizializzazione modulo di scansione...")
		}
		if targetIP == "" {
			logger.Log.Info("Errore: devi specificare un target. Usa --help per info.")
			return
		}
		logger.Log.Infof("Avviata scansione sul target: %s", targetIP)

		timeout, _ := cobraCmd.Flags().GetDuration("timeout")
		if timeout == 0 {
			timeout = config.Config.Options.Timeout
		}
		logger.Log.Infof("Timeout usato: %s", timeout)

		open := ScanPort(targetIP, targetPort, timeout)
		if open {
			logger.Log.Infof("✅ Porta %d aperta su %s", targetPort, targetIP)
		} else {
			logger.Log.Infof("❌ Porta %d chiusa o non raggiungibile su %s", targetPort, targetIP)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(scanCmd)
	scanCmd.Flags().Duration("timeout", 0, "Timeout connessione (es: 500ms, 1s)")
	scanCmd.Flags().StringVarP(&targetIP, "target", "t", "", "IP di destinazione da scansionare")
	scanCmd.Flags().IntVarP(&targetPort, "port", "p", 80, "Porta da scansionare")
}

func ScanPort(host string, port int, timeout time.Duration) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		logger.Log.Errorf("porta %d chiusa su %s: %v", port, host, err)
		return false
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			logger.Log.Errorf("porta %d chiusa su %s: %v", port, host, err)
		}
	}(conn)
	return true
}
