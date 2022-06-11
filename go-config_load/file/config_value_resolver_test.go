package file

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var _ Config = (*mockConfig)(nil)

type mockConfig struct {
	mock.Mock
}

func (_m *mockConfig) GetString(key string) string {
	ret := _m.Called(key)
	return ret.Get(0).(string)
}

func (_m *mockConfig) GetInt(key string) (int, error) {
	ret := _m.Called(key)
	return ret.Get(0).(int), ret.Error(1)
}

const testConfKey = "TEST_CONFIG_KEY"

func TestResolveValueStringOrDefault(t *testing.T) {
	var cases = map[string]struct {
		isNotEmptySysConf bool
		inputSysConfValue string
		isNotEmptyUsrConf bool
		inputUsrConfValue string
		expected          string
	}{
		"system_exist_user_exist":       {true, "system", true, "user", "user"},
		"system_exist_user_nonexist":    {true, "system", false, "", "system"},
		"system_nonexist_user_exist":    {false, "", true, "user", "system"},
		"system_nonexist_user_nonexist": {false, "", false, "", "default"},
	}

	buildMockConf := func(inputConfValue string) Config {
		mockConf := new(mockConfig)
		mockConf.On("GetString", testConfKey).Return(inputConfValue)
		return mockConf
	}

	for testCaseName, tt := range cases {
		t.Run(testCaseName, func(t *testing.T) {
			t.Parallel()
			var sysConf Config = EmptyConfig{}
			if tt.isNotEmptySysConf {
				sysConf = buildMockConf(tt.inputSysConfValue)
			}
			var usrConf Config = EmptyConfig{}
			if tt.isNotEmptyUsrConf {
				usrConf = buildMockConf(tt.inputUsrConfValue)
			}
			resolver := NewConfigValueResolver(sysConf, usrConf)
			assert.Equal(t, tt.expected, resolver.ResolveValueStringOrDefault(testConfKey, "default"))
		})
	}
}

func TestResolveValueIntOrDefault(t *testing.T) {
	defaultValue := 0
	invalidValue := "invalid"
	nonExistKey := ""
	var cases = map[string]struct {
		isNotEmptySysConf bool
		inputSysConfValue string
		isNotEmptyUsrConf bool
		inputUsrConfValue string
		expected          int
	}{
		"system_exist_user_exist":       {true, "10", true, "20", 20},
		"system_exist_user_nonKey":      {true, "10", true, nonExistKey, 10},
		"system_exist_user_invalid":     {true, "10", true, invalidValue, 10},
		"system_noKey_user_exist":       {true, nonExistKey, true, "20", 20},
		"system_invalid_user_exist":     {true, invalidValue, true, "20", 20},
		"system_nonexist_user_exist":    {false, "", true, "20", 20},
		"system_nonexist_user_nonKey":   {false, "", true, nonExistKey, defaultValue},
		"system_nonexist_user_invalid":  {false, "", true, invalidValue, defaultValue},
		"system_exist_user_nonexist":    {true, "10", false, "", 10},
		"system_nonKey_user_nonexist":   {true, nonExistKey, false, "", defaultValue},
		"system_invalid_user_nonexist":  {true, invalidValue, false, "", defaultValue},
		"system_nonexist_user_nonexist": {false, "", false, "", defaultValue},
	}

	buildMockConf := func(inputConfValue string) Config {
		mockConf := new(mockConfig)
		mockConf.On("GetInt", testConfKey).Return(strconv.Atoi(inputConfValue))
		return mockConf
	}

	for testCaseName, tt := range cases {
		t.Run(testCaseName, func(t *testing.T) {
			//t.Parallel()
			var sysConf Config = EmptyConfig{}
			if tt.isNotEmptySysConf {
				sysConf = buildMockConf(tt.inputSysConfValue)
			}
			var usrConf Config = EmptyConfig{}
			if tt.isNotEmptyUsrConf {
				usrConf = buildMockConf(tt.inputUsrConfValue)
			}
			resolver := NewConfigValueResolver(sysConf, usrConf)
			val := resolver.ResolveValueIntOrDefault(testConfKey, defaultValue)
			assert.Equal(t, tt.expected, val)
		})
	}
}
