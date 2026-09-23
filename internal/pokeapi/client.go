package pokeapi

import (
	"net/http"
	"time"

	"github.com/danielakinremi1-dev/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	pokeCache  *pokecache.Cache
}

func NewClient(timeout time.Duration, cacheReset time.Duration) Client {
	return Client{
		httpClient: http.Client{Timeout: timeout},
		pokeCache:  pokecache.NewCache(cacheReset),
	}
}
