package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

type NetklitFormatter struct{}

func (f *NetklitFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.InfoLevel:
		levelColor = 32 // Verde
	case logrus.WarnLevel:
		levelColor = 33 // Giallo
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = 31 // Rosso
	default:
		levelColor = 36 // Cyan
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	levelName := strings.ToUpper(entry.Level.String())

	output := fmt.Sprintf("[%s] [\x1b[%dm%s\x1b[0m] %s\n",
		timestamp,
		levelColor,
		levelName,
		entry.Message)

	return []byte(output), nil
}

func Init() {
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		DisableColors: false, //  true per log in file
	})
	Log.SetOutput(os.Stdout)
	Log.SetFormatter(&NetklitFormatter{})

	lvl, err := logrus.ParseLevel("debug")
	if err != nil {
		Log.Warnf("Error parsing log level: %v setting default to INFO", err)
		lvl = logrus.InfoLevel
	}
	Log.SetLevel(lvl)

}
