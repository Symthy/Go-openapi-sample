package flags

type FeatureConfig interface {
	// プロジェクト名 + ONでフラグ名とする
	ProjectAOn() bool
	ProjectBOn(userID uint64) bool
}

// 設定ファイルを元に環境変数が設定されているので、環境変数からフラグ設定を読み取る
type EnvConfig struct {
	ProjectAOn     bool `required:"true" envconfig:"ProjectAOn" default:"false"`
	ProjectBOn     bool `required:"true" envconfig:"ProjectBOn" default:"false"`
	WhitelistUsers []uint64
}

type featureConfig struct {
	envConfig EnvConfig
}

func (f featureConfig) ProjectAOn() bool {
	return f.envConfig.ProjectAOn
}

func (f featureConfig) ProjectBOn(userID uint64) bool {
	//whitelistUserIDs := []uint64{100, 111}
	for _, wuid := range f.envConfig.WhitelistUsers {
		if wuid == userID {
			return true
		}
	}
	return f.envConfig.ProjectBOn
}

func NewFeatureConfig(envConfig EnvConfig) FeatureConfig {
	return &featureConfig{envConfig: envConfig}
}
