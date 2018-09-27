package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgLink struct {
	Address sdk.AccAddress `json:"address"`
	ContentID1 string `json:"cid1"`
	ContentID2 string `json:"cid2"`
}

var _ sdk.Msg = MsgLink{}

func NewMsgLink(address sdk.AccAddress, cid1 string, cid2 string) MsgLink {
	return MsgLink{Address: address, ContentID1: cid1, ContentID2: cid2}
}

func (MsgLink) Type() string { return "link" }

func (msg MsgLink) ValidateBasic() sdk.Error {

	if len(msg.Address) == 0 {
		return sdk.ErrInvalidAddress(msg.Address.String())
	}

	if len(msg.ContentID1) == 0 || len(msg.ContentID2) == 0 {
		return ErrInvalidCid(DefaultCodespace).TraceSDK("")
	}

	return nil
}

func (msg MsgLink) GetSignBytes() []byte {
	b, err := msgCdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgLink) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}