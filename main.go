package main

import (
	router "magiot-game-api/api/router"
)

func main() {
	r := router.NewRouter()
	r.Logger.Fatal(r.Start(":3000"))
}
