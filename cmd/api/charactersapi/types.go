package charactersapi

type Character struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Debut       int    `json:"debut"`
}

type Offset struct {
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}

type CharacterList struct {
	Offset
	Characters []Character `json:"characters"`
}
