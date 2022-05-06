package conf

type Config struct {
	DB *DBConfig
}

//Config Database Postgresql
type DBConfig struct {
	Username string
	Password string
	Dbname   string
	Host     string
	Port     string
}

//Set configs
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
