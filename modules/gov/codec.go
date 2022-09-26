package gov

import (
	spartangov "github.com/bianjieai/spartan-cosmos/module/gov/module"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		gov.AppModuleBasic{},
		spartangov.AppModuleBasic{},
	)
}
