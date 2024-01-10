package entity

type Pet struct {
	Breed  string `json:"breed"`
	Weight Weight `json:"weight"`
}

type Weight struct {
	Value   float64 `json:"value"`
	Metrics string  `json:"metrics"`
}
