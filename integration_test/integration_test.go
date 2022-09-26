package integration

import (
	spartanchain_mod_parser "github.com/kaifei-bianjie/spartanchain-mod-parser"
	"github.com/stretchr/testify/suite"
	"testing"
)

type IntegrationTestSuite struct {
	spartanchain_mod_parser.MsgClient
	suite.Suite
}

type SubTest struct {
	testName string
	testCase func(s IntegrationTestSuite)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.MsgClient = spartanchain_mod_parser.NewMsgClient()
}
