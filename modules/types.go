package msgs

import (
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	//gov
	MsgTypeSubmitProposal = "submit_proposal"
	MsgTypeDeposit        = "deposit"
	MsgTypeVote           = "vote"
	MsgTypeVoteWeighted   = "vote_weighted"
)

type (
	//gov
	MsgDeposit        = gov.MsgDeposit
	MsgSubmitProposal = gov.MsgSubmitProposal
	TextProposal      = gov.TextProposal
	MsgVote           = gov.MsgVote
	Proposal          = gov.Proposal
	SdkVote           = gov.Vote
	GovContent        = gov.Content
	MsgVoteWeighted   = gov.MsgVoteWeighted
)
