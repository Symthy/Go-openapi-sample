package file

type NonExistKeyError struct{}

func (e NonExistKeyError) Error() string {
	return "non exist key"
}

type NoValueError struct{}

func (e NoValueError) Error() string {
	return "non value"
}
