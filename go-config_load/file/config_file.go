package file

import (
	"fmt"

	"gopkg.in/go-ini/ini.v1"
)

type ConfFile interface {
	GetString(key string) string
	GetInt(key string) (int, error)
}

type ConfFileLoader struct {
	file *ini.File
	nonExist bool
}

func NewConfFileLoader(filePath string) ConfFile {
	cfg, err := ini.Load(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &ConfFileLoader{file: cfg}
}

func (c ConfFileLoader) GetString(key string) string {
	return c.file.Section("").Key(key).String()
}

func (c ConfFileLoader) GetInt(key string) (int, error) {
	return c.file.Section("").Key(key).Int()
}