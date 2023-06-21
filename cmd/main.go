package main

import (
	"os"

	"github.com/NicholeMattera/Harmony/internal/api"
)

func main() {
	r := api.SetupRouter()

	if addr, addrExists := os.LookupEnv("HARMONY_LISTEN_ADDR"); !addrExists {
		r.Run("0.0.0.0:8080")
	} else {
		r.Run(addr)
	}
}
