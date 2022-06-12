package config

// Databases

type Redis struct {
	Address string `env:"REDIS_ADDRESS" env-default:"0.0.0.0:6379"`
	Pass    string `env:"REDIS_PASS"    env-default:""`
	DB      int    `env:"REDIS_DB"      env-default:"0"`
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
	Redis Redis
	//
	Host Host
}
