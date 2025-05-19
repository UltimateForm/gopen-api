package config

type DbConfig struct {
	Uri      string
	Username string
	Password string
}

type AuthConfig struct {
	JwtSign string
}
