package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/Symthy/golang-practices/go-godoenv/file"
)

func main() {
	sysConf := file.LoadConfigFile("./conf/system.propertie")
	usrConf := file.LoadConfigFile("./conf/user.propertie")
	resolver := file.NewConfigValueResolver(sysConf, usrConf)

	fmt.Println(resolver.ResolveValueStringOrDefault("ENV", "default"))
	fmt.Println(resolver.ResolveValueStringOrDefault("DB", "default"))
	fmt.Println(resolver.ResolveValueStringOrDefault("DB_USER", "default"))
	fmt.Println(resolver.ResolveValueStringOrDefault("DB_PASS", "default"))
}

func resolveValue(conf file.Config, key string) string {
	if conf == nil {
		fmt.Println("no conf file")
		return ""
	}
	return conf.GetString(key)
}

// OSS使わず、<key>=<value> を読み込む
func LoadProperties(filePath string) (map[string]string, error) {
	keyToValue := map[string]string{}
	lines, err := readAllLines(filePath)
	if err != nil {
		return nil, err
	}
	commentLineRegx := regexp.MustCompile(`^[ ]*#.*`)
	keyValueLineRegx := regexp.MustCompile(`^[ ]*([^ ]*)[ ]*=[ ]*([^ ]*)[ ]*$`)
	for _, line := range lines {
		if commentLineRegx.MatchString(line) {
			continue
		}
		if keyValueLineRegx.MatchString(line) {
			group := keyValueLineRegx.FindAllStringSubmatch(line, -1)
			key := group[0][1]
			value := group[0][2]
			keyToValue[key] = value
		}
	}
	return keyToValue, nil
}

func readAllLines(filePath string) ([]string, error) {
	if _, err := os.Stat(filePath); err != nil {
		// non exist file
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, err
}
