package config

// Config : structure of the YAML file configuration
type Config struct {
	Server   serverConfig   `yaml:"server"`
	Database databaseConfig `yaml:"database"`
	Logger   loggerConfig   `yaml:"logger"`
}

// ServerConfig : structure for server config
type serverConfig struct {
	Port   string `yaml:"port"`
	Host   string `yaml:"host"`
	Secret string `yaml:"secret"`
}

// DatabaseConfig : structure for database config
type databaseConfig struct {
	Host         string `yaml:"host"`
	DatabaseName string `yaml:"dbName"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
}

type loggerConfig struct {
	Path     string `yaml:"path"`
	FileName string `yaml:"filename"`
}
