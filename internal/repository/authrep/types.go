package authrep

type User struct {
	Email    string
	Password string
}

func (src User) ToParams() map[string]any {
	return map[string]any{
		"email":    src.Email,
		"password": src.Password,
	}
}
