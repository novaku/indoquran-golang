package config

// Config : structure of the YAML file configuration
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Logger   LoggerConfig   `yaml:"logger"`
	Redis    RedisConfig    `yaml:"redis"`
}

// ServerConfig : structure for server config
type ServerConfig struct {
	Port   string `yaml:"port"`
	Host   string `yaml:"host"`
	Secret string `yaml:"secret"`
}

// DatabaseConfig : structure for database config
type DatabaseConfig struct {
	Host         string `yaml:"host"`
	DatabaseName string `yaml:"dbName"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
}

// LoggerConfig : configuration for logger
type LoggerConfig struct {
	Path     string `yaml:"path"`
	FileName string `yaml:"filename"`
}

// RedisConfig : configuration fo redis
type RedisConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}
