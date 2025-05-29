package charrep

type Offset struct {
	Limit int
	Skip  int
}

func (src Offset) ToParams() map[string]any {
	return map[string]any{
		"limit": src.Limit,
		"skip":  src.Skip,
	}
}

type Character struct {
	Id          string
	Name        string
	Description string
	Debut       int
}
