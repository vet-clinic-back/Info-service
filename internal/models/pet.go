package models

type Pet struct {
	ID             uint    `json:"id"`
	AnimalType     string  `json:"animal_type"`
	Name           string  `json:"name"`
	Gender         string  `json:"gender,omitempty"`
	Age            uint    `json:"age,omitempty"`
	Weight         float64 `json:"weight,omitempty"`
	Condition      string  `json:"condition,omitempty"`
	Behavior       string  `json:"behavior,omitempty"`
	ResearchStatus string  `json:"research_status,omitempty"`
}
