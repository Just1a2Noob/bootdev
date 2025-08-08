package pokeapi

type Pokemon struct {
	Name           string `json:"name,omitempty"`
	BaseExperience int    `json:"base_experience,omitempty"`
}

type Pokedex struct {
	pokedex map[string]Pokemon
}
