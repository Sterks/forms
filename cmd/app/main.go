package main

import "forms/internal/app"

const configsDir = "configs"

func main() {

	// readConfig()
	app.Run(configsDir)
}
