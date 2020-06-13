package config

// Config : structure of the YAML file configuration
type Config struct {
	Server   serverConfig   `yaml:"server"`
	Database databaseConfig `yaml:"database"`
	Redis    redisConfig    `yaml:"redis"`
}

type serverConfig struct {
	Port   string `yaml:"port"`
	Host   string `yaml:"host"`
	Secret string `yaml:"secret"`
}

type databaseConfig struct {
	Host         string `yaml:"host"`
	DatabaseName string `yaml:"dbName"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
}

type redisConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}
