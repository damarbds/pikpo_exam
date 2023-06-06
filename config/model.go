package config

const (
	confPathPrefix           = "./configs/"
	confPathSuffix           = "_config.json"
	runtimeEnvironmentEnvKey = "APP_ENV"
)

type confData struct {
	Server   serverConf   `json:"server,omitempty"`
	Database DatabaseConf `json:"database,omitempty"`
}

type DatabaseConf struct {
	Postgres PostgresConf `json:"postgres,omitempty"`
	Type     string       `json:"type"`
}

type PostgresConf struct {
	ConnectionString string `json:"connectionString"`
}

type serverConf struct {
	ListenPort string `json:"listenPort,omitempty"`
}
