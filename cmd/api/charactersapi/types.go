package charactersapi

import "github.com/UltimateForm/gopen-api/internal/core"

type CharacterInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Debut       int    `json:"debut"`
}

func (src CharacterInfo) Validate() error {
	if src.Name == "" || src.Description == "" || src.Debut == 0 {
		return core.BadRequest("missing one or more parameters: name, description, debut")
	}
	return nil
}

type Character struct {
	Id string `json:"id"`
	CharacterInfo
}

type Offset struct {
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}

type CharacterList struct {
	Offset
	Characters []Character `json:"characters"`
}
