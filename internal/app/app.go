package app

import "github.com/NicholeMattera/Harmony/internal/store"

type AppLayer interface {}

type appLayer struct {
    store store.StoreLayer
}

func New(storeLayer store.StoreLayer) *appLayer {
    return &appLayer{
        store: storeLayer,
    }
}
