package config

type Cfg struct {
	Service string `toml:"service"`
	Server  string `toml:"server"`
	Proxy   string `toml:"proxy"`
	Web     string `toml:"web"`
	App     struct {
		Name string `toml:"name"`
	} `toml:"app"`
	Log struct {
		LogLevel string `toml:"log_level"`
	} `toml:"log"`
	Db []struct {
		Driver string `toml:"driver"`
		Dburl  string `toml:"dburl"`
	} `toml:"db"`
}
