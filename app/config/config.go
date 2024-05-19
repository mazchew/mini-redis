package config

type Config struct {
	Dir string
	DbFilename string
}

func NewConfig(dir string, dbFilename string) *Config {
	return &Config{
		Dir: dir,
		DbFilename: dbFilename,
	}
}
