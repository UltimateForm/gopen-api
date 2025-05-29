package charactersapi

type Character struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Debut       int    `json:"debut"`
}

type CharacterList struct {
	Offset
	Characters []Character `json:"characters"`
}
