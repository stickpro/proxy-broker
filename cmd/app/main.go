package main

import "asocks-ws/internal/app"

const configDir = "configs"

func main() {
	app.Run(configDir)
}
