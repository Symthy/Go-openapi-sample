package flags

import "fmt"

type Charge interface {
	Charge(userID uint64, amount uint64)
}

type chargeService struct {
	featureConfig FeatureConfig
}

func (s *chargeService) Charge(userID uint64, amount uint64) {
	if !s.featureConfig.ProjectBOn(userID) {
		// 既存処理
		fmt.Println("既存処理")
	} else {
		// 追加処理
		fmt.Println("改修後処理")
	}

	// charge実行
}

// DIツールで依存注入。テストではfeatureConfigのmockを使えばいい
func NewChargeService(featureConfig FeatureConfig) Charge {
	return &chargeService{featureConfig: featureConfig}
}
