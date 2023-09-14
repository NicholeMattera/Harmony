package main

import (
	"github.com/NicholeMattera/Harmony/internal/app"
	"github.com/NicholeMattera/Harmony/internal/http"
	"github.com/NicholeMattera/Harmony/internal/store"
)

func main() {
    storeLayer := store.New()
	appLayer := app.New(storeLayer)
	httpLayer := http.New(appLayer)

	httpLayer.Run()
}
