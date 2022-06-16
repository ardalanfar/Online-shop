package config

//config system
type Config struct {
	DB    *DBConfig
	Email *SendEmail
}

//config database postgresql
type DBConfig struct {
	Username string
	Password string
	Dbname   string
	Host     string
	Port     string
}

//email config
type SendEmail struct {
	From     string
	Username string
	Password string
	SmtpHost string
	SmtpPort string
}

//set config app
func GetConfig() *Config {
	return &Config{

		DB: &DBConfig{
			Username: "admin",
			Password: "123456",
			Dbname:   "farashop",
			Host:     "localhost",
			Port:     "9920",
		},

		Email: &SendEmail{
			//Sender data.
			From:     "",
			Username: "",
			Password: "",
			//smtp server configuration.
			SmtpHost: "smtp.gmail.com",
			SmtpPort: "587",
		},
	}
}
