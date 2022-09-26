package main

import "github.com/Symthy/golang-practices/go-feature-flags/internal/flags"

func main() {
	conf := flags.EnvConfig{
		ProjectAOn:     true,
		ProjectBOn:     false,
		WhitelistUsers: []uint64{100, 111},
	}

	featureConf := flags.NewFeatureConfig(conf)

	serv := flags.NewChargeService(featureConf)
	serv.Charge(101, 1000)
}
