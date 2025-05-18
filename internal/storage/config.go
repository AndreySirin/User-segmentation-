package storage

type Config struct {
	DbName   string `yaml:"dbName"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
}
