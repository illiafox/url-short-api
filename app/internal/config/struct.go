package config

// Databases

type Postgres struct {
	User     string `env:"POSTGRES_USER"     env-required:""`
	Pass     string `env:"POSTGRES_PASSWORD" env-required:""`
	DBName   string `env:"POSTGRES_DB"       env-required:""`
	IP       string `env:"POSTGRES_IP"       env-required:""`
	Port     int    `env:"POSTGRES_PORT"     env-default:"5432"`
	Protocol string `env:"POSTGRES_PROTOCOL" env-default:"tcp"`
}

// Host

type Host struct {
	Addr string `env:"HOST_ADDRESS"   env-default:"0.0.0.0"`
	Port int    `env:"HOST_PORT"      env-default:"8080"`
	Key  string `env:"HOST_KEY_PATH"`  // Path to TLS key
	Cert string `env:"HOST_CERT_PATH"` // Path to TLS certificate
}

//

type Config struct {
	Postgres Postgres
	//
	Host Host
}
