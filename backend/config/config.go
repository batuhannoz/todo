package config

type Config struct {
	*MySQL
	*Server
}

type MySQL struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Server struct {
}

func GetConfig(path string) *Config {

	return &Config{
		&MySQL{},
		&Server{},
	}
}
