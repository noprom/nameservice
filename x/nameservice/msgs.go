package nameservice

import (
  "encoding/json"
  sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgSetName defines a SetName message
type MsgSetName struct {
	NameID string
	Value  string
	Owner  sdk.AccAddress
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
  return MsgSetName{
    NameID: name,
    Value: value,
    Owner: owner
  }
}

// Name should return the action
func (msg MsgSetName) Route() string { return "nameservice" }

// Type should return the name of the module
func (msg MsgSetName) Type() string { return "set_name" }

// ValdateBasic Implements Msg.
func (msg MsgSetName) ValdateBasic() sdk.Error {
  if msg.Owner.Empty() {
    return sdk.ErrInvalidAddress(msg.Owner.string())
  }

  if (len(msg.NameID) == 0 || len(msg.Value) == 0) {
    return sdk.ErrUnknownRequest("Name and/or Value cannot be empty")
  }

  return nil
}

// GetSignBytes Implements Msg.
func (msg MsgSetName) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}