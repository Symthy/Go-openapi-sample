package file

type ConfigValueResolver struct {
	systemConf ConfFile
	userConf   ConfFile
}

func NewConfigValueResolver(systemConf ConfFile, userConf ConfFile) ConfigValueResolver {
	return ConfigValueResolver{systemConf: systemConf, userConf: userConf}
}

func (c ConfigValueResolver) ResolveValueString(key string, defaultValue string) string {
	value := defaultValue
	value = getConfValueStringOrElse(c.systemConf, key, value)
	value = getConfValueStringOrElse(c.userConf, key, value)
	return value
}

func (c ConfigValueResolver) ResolveValueInt(key string, defaultValue int) int {
	value := defaultValue
	value = getConfValueIntOrElse(c.systemConf, key, value)
	value = getConfValueIntOrElse(c.userConf, key, value)
	return value
}

func getConfValueStringOrElse(conf ConfFile, key string, value string) string {
	if conf == nil {
		return value
	}
	if confValue := conf.GetString(key); confValue != "" {
		return confValue
	}
	return value
}

func getConfValueIntOrElse(conf ConfFile, key string, value int) int {
	if conf == nil {
		return value
	}
	if confValue, err := conf.GetInt(key); err == nil {
		return confValue
	}
	return value
}
