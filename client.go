package spartanchain_mod_parser

import (
	. "github.com/kaifei-bianjie/common-parser"
	"github.com/kaifei-bianjie/spartanchain-mod-parser/codec"
	"github.com/kaifei-bianjie/spartanchain-mod-parser/modules/gov"
)

type MsgClient struct {
	Gov Client
}

func NewMsgClient() MsgClient {
	codec.MakeEncodingConfig()
	return MsgClient{
		Gov: gov.NewClient(),
	}
}
