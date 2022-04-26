package config

//config database mysql
type Postgresql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}
