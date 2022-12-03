package config

type Config struct {
	Logger   *Logger   `mapstructure:"logger" yaml:"logger"`
	Database *Database `mapstucture:"database" yaml:"database"`
	Server   *Server   `mapstructure:"server" yaml:"server"`
}
