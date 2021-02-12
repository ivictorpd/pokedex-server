package models

type Pokemon []struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Classification string   `json:"classification"`
	Types          []string `json:"types"`
	Resistant      []string `json:"resistant"`
	Weaknesses     []string `json:"weaknesses"`
	IsFavorite     bool     `json:"isfavorite"`
	ImageUrl       string   `json:"imageurl"`
	Sound          string   `json:"sound"`
	Weight         struct {
		Minimum string `json:"minimum"`
		Maximum string `json:"maximum"`
	} `json:"weight"`
	Height struct {
		Minimum string `json:"minimum"`
		Maximum string `json:"maximum"`
	} `json:"height"`
	FleeRate              float64 `json:"fleeRate"`
	EvolutionRequirements struct {
		Amount int    `json:"amount"`
		Name   string `json:"name"`
	} `json:"evolutionRequirements,omitempty"`
	Evolutions []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"evolutions,omitempty"`
	MaxCP   int `json:"maxCP"`
	MaxHP   int `json:"maxHP"`
	Attacks struct {
		Fast []struct {
			Name   string `json:"name"`
			Type   string `json:"type"`
			Damage int    `json:"damage"`
		} `json:"fast"`
		Special []struct {
			Name   string `json:"name"`
			Type   string `json:"type"`
			Damage int    `json:"damage"`
		} `json:"special"`
	} `json:"attacks"`
	PreviousEvolutionS []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"Previous evolution(s),omitempty"`
	CommonCaptureArea   string `json:"Common Capture Area,omitempty"`
	Asia                string `json:"Asia,omitempty"`
	AustraliaNewZealand string `json:"Australia, New Zealand,omitempty"`
	WesternEurope       string `json:"Western Europe,omitempty"`
	NorthAmerica        string `json:"North America,omitempty"`
	PokMonClass         string `json:"Pokémon Class,omitempty"`
	LEGENDARY           string `json:"LEGENDARY,omitempty"`
	MYTHIC              string `json:"MYTHIC,omitempty"`
}
