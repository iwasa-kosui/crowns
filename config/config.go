package config

type NetworkConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type ServerConfig struct {
	NetworkConfig
	DatabaseConfig
}

func Load() *ServerConfig {
	return &ServerConfig{
		NetworkConfig{
			Port: "8081",
		},
		DatabaseConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "crowns",
			Password: "Cr0wnS",
			DbName:   "crowns",
		},
	}
}
