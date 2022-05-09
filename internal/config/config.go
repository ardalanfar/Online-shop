package config

//config system
type Config struct {
	DB *DBConfig
}

//config database postgresql
type DBConfig struct {
	Username string
	Password string
	Dbname   string
	Host     string
	Port     string
}

//set configs
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Username: "admin",
			Password: "123456",
			Dbname:   "farashop",
			Host:     "localhost",
			Port:     "9920",
		},
	}
}
