package pokeapi

import (
	"net/http"
	"time"

	"github.com/maximilian-1011/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
	dex        Dex
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
		dex: NewDex(),
	}
}
