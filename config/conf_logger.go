package config

type Logger struct {
	Level        string `yaml:"host"`
	Prefix       int    `yaml:"port"`
	Director     string `yaml:"user"`
	ShowLine     bool   `yaml:"database"`
	LogInConsole bool   `yaml:"database"`
}
