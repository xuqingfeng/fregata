package logging

type Config struct {
	File  string `toml:"file"`
	Level string `toml:"level"`
}

func NewConfig() Config {

	return Config{}
}
