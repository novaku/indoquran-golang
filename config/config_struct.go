package config

// Config : structure of the YAML file configuration
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

// ServerConfig : structure for server config
type ServerConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

// DatabaseConfig : structure for database config
type DatabaseConfig struct {
	Host         string `yaml:"host"`
	DatabaseName string `yaml:"dbName"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
}
