package main

import (
	"netklit/cmd"
	_ "netklit/cmd/scan"
	"netklit/internal/config"
	"netklit/pkg/logger"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	logger.Init()
	config.Execute()
	cmd.Execute()
}
