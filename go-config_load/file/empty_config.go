package file

var _ Config = (*EmptyConfig)(nil)

type EmptyConfig struct{}

func (EmptyConfig) GetString(ket string) string {
	return ""
}

func (EmptyConfig) GetInt(key string) (int, error) {
	return 0, NoValueError{}
}
