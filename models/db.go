package models

type Db struct {
	Type     string `toml:"type"`
	Addr     string `toml:"addr"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"database"`
	Charset  string `toml:"charset"`
}
