package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
	"sync"
)

type Dex struct {
	list map[string]Pokemon
	mu   *sync.RWMutex
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}
	if err := json.Unmarshal(data, &pokemonResp); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	c.dex.Add(pokemonResp)
	return pokemonResp, nil
}

func NewDex() Dex {
	d := Dex{
		list: make(map[string]Pokemon),
		mu:   &sync.RWMutex{},
	}

	return d
}

func (d *Dex) Add(pokemon Pokemon) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.list[pokemon.Name] = pokemon
}

func (c *Client) DexGet(name string) (Pokemon, bool) {
	c.dex.mu.RLock()
	defer c.dex.mu.RUnlock()
	val, ok := c.dex.list[name]
	return val, ok
}

func (c *Client) GetDex() map[string]Pokemon {
	return c.dex.list
}