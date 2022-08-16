package config

type DbDriverStruct struct {
	Drive  string
	Entity DbInfo
}

type DbInfo struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DbName   string `toml:"dbname"`
}
