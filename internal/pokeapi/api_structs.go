package pokeapi

type LocationArea struct {
	Name string
	URL  string
}

type PokeAPILocationResponse struct {
	Count    int
	Next     *string
	Previous *string
	Results  []LocationArea
}
