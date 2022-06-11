package file

type ConfigValueResolver struct {
	systemConf Config
	userConf   Config
}

func NewConfigValueResolver(systemConf Config, userConf Config) ConfigValueResolver {
	return ConfigValueResolver{systemConf: systemConf, userConf: userConf}
}

func (c ConfigValueResolver) ResolveValueStringOrDefault(key string, defaultValue string) string {
	value := defaultValue
	value = getConfValueStringOrElse(c.systemConf, key, value)
	value = getConfValueStringOrElse(c.userConf, key, value)
	return value
}

func (c ConfigValueResolver) ResolveValueIntOrDefault(key string, defaultValue int) int {
	value := defaultValue
	value = getConfValueIntOrElse(c.systemConf, key, value)
	value = getConfValueIntOrElse(c.userConf, key, value)
	return value
}

func getConfValueStringOrElse(conf Config, key string, defaultValue string) string {
	if confValue := conf.GetString(key); confValue != "" {
		return confValue
	}
	return defaultValue
}

func getConfValueIntOrElse(conf Config, key string, defaultValue int) int {
	if confValue, err := conf.GetInt(key); err == nil {
		return confValue
	}
	return defaultValue
}
