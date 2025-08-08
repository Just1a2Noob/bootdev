package pokeapi

type Pokemon struct {
	Name           string `json:"name,omitempty"`
	BaseExperience int    `json:"base_experience,omitempty"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
	Height int `json:"height"`
}
