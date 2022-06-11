package file

import (
	"gopkg.in/go-ini/ini.v1"
)

var _ Config = (*ConfigFile)(nil)

type Config interface {
	GetString(key string) string
	GetInt(key string) (int, error)
}

func LoadConfigFile(filePath string) Config {
	cfg, err := ini.Load(filePath)
	if err != nil {
		return &EmptyConfig{}
	}
	return &ConfigFile{file: cfg}
}

type ConfigFile struct {
	file *ini.File
}

func (c ConfigFile) GetString(key string) string {
	return c.file.Section("").Key(key).String()
}

func (c ConfigFile) GetInt(key string) (int, error) {
	if c.file.Section("").HasKey(key) {
		return c.file.Section("").Key(key).Int()
	}
	return 0, NonExistKeyError{}
}
