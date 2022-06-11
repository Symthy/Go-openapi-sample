package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessToGetString(t *testing.T) {
	var cases = []struct {
		inputKey string
		expected string
	}{
		{
			inputKey: "TEST_FOR_GET_STRING",
			expected: "test",
		},
	}
	config := LoadConfigFile("../test/conf/test.conf")
	for _, tt := range cases {
		assert.Equal(t, tt.expected, config.GetString(tt.inputKey))
	}
}

func TestSuccessToGetInt(t *testing.T) {
	var cases = []struct {
		inputKey string
		expected int
	}{
		{
			inputKey: "TEST_FOR_GET_INT_SUCCESS",
			expected: 10,
		},
	}
	config := LoadConfigFile("../test/conf/test.conf")
	for _, tt := range cases {
		val, err := config.GetInt(tt.inputKey)
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		assert.Equal(t, tt.expected, val)
	}
}

func TestFailedToGetInt(t *testing.T) {
	var cases = map[string]struct {
		inputKey            string
		expectedErrorString string
	}{
		"parse_err":     {"TEST_FOR_GET_INT_ERROR", "strconv.ParseInt: parsing \"string\": invalid syntax"},
		"non_exist_key": {"NON_EXIST_KEY", NonExistKeyError{}.Error()},
	}
	config := LoadConfigFile("../test/conf/test.conf")
	for _, tt := range cases {
		_, err := config.GetInt(tt.inputKey)
		assert.EqualError(t, err, tt.expectedErrorString)
	}
}
