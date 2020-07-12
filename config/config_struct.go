package config

// Config : structure of the YAML file configuration
type Config struct {
	Server   serverConfig   `yaml:"server"`
	Auth     authConfig     `yaml:"auth"`
	Database databaseConfig `yaml:"database"`
	Redis    redisConfig    `yaml:"redis"`
}

type serverConfig struct {
	Port   string `yaml:"port"`
	Host   string `yaml:"host"`
	Secret string `yaml:"secret"`
}

type authConfig struct {
	AccessSecret  string `yaml:"access_secret"`
	RefreshSecret string `yaml:"refresh_secret"`
	AccessExpire  int    `yaml:"access_expire"`
	RefreshExpire int    `yaml:"refresh_expire"`
}

type databaseConfig struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	DatabaseName string `yaml:"dbName"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
}

type redisConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}
