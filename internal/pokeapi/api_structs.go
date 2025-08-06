package pokeapi

type LocationArea struct {
	Name string
	URL  string
}

type PokeAPIPageResponse[T any] struct {
	Count    int
	Next     *string
	Previous *string
	Results  []T
}
