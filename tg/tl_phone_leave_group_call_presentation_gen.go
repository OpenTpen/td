// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// PhoneLeaveGroupCallPresentationRequest represents TL type `phone.leaveGroupCallPresentation#1c50d144`.
// Stop screen sharing in a group call
//
// See https://core.telegram.org/method/phone.leaveGroupCallPresentation for reference.
type PhoneLeaveGroupCallPresentationRequest struct {
	// The group call
	Call InputGroupCall
}

// PhoneLeaveGroupCallPresentationRequestTypeID is TL type id of PhoneLeaveGroupCallPresentationRequest.
const PhoneLeaveGroupCallPresentationRequestTypeID = 0x1c50d144

// Ensuring interfaces in compile-time for PhoneLeaveGroupCallPresentationRequest.
var (
	_ bin.Encoder     = &PhoneLeaveGroupCallPresentationRequest{}
	_ bin.Decoder     = &PhoneLeaveGroupCallPresentationRequest{}
	_ bin.BareEncoder = &PhoneLeaveGroupCallPresentationRequest{}
	_ bin.BareDecoder = &PhoneLeaveGroupCallPresentationRequest{}
)

func (l *PhoneLeaveGroupCallPresentationRequest) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.Call.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *PhoneLeaveGroupCallPresentationRequest) String() string {
	if l == nil {
		return "PhoneLeaveGroupCallPresentationRequest(nil)"
	}
	type Alias PhoneLeaveGroupCallPresentationRequest
	return fmt.Sprintf("PhoneLeaveGroupCallPresentationRequest%+v", Alias(*l))
}

// FillFrom fills PhoneLeaveGroupCallPresentationRequest from given interface.
func (l *PhoneLeaveGroupCallPresentationRequest) FillFrom(from interface {
	GetCall() (value InputGroupCall)
}) {
	l.Call = from.GetCall()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneLeaveGroupCallPresentationRequest) TypeID() uint32 {
	return PhoneLeaveGroupCallPresentationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneLeaveGroupCallPresentationRequest) TypeName() string {
	return "phone.leaveGroupCallPresentation"
}

// TypeInfo returns info about TL type.
func (l *PhoneLeaveGroupCallPresentationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.leaveGroupCallPresentation",
		ID:   PhoneLeaveGroupCallPresentationRequestTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Call",
			SchemaName: "call",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *PhoneLeaveGroupCallPresentationRequest) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode phone.leaveGroupCallPresentation#1c50d144 as nil")
	}
	b.PutID(PhoneLeaveGroupCallPresentationRequestTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *PhoneLeaveGroupCallPresentationRequest) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode phone.leaveGroupCallPresentation#1c50d144 as nil")
	}
	if err := l.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.leaveGroupCallPresentation#1c50d144: field call: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (l *PhoneLeaveGroupCallPresentationRequest) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode phone.leaveGroupCallPresentation#1c50d144 to nil")
	}
	if err := b.ConsumeID(PhoneLeaveGroupCallPresentationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.leaveGroupCallPresentation#1c50d144: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *PhoneLeaveGroupCallPresentationRequest) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode phone.leaveGroupCallPresentation#1c50d144 to nil")
	}
	{
		if err := l.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.leaveGroupCallPresentation#1c50d144: field call: %w", err)
		}
	}
	return nil
}

// GetCall returns value of Call field.
func (l *PhoneLeaveGroupCallPresentationRequest) GetCall() (value InputGroupCall) {
	if l == nil {
		return
	}
	return l.Call
}

// PhoneLeaveGroupCallPresentation invokes method phone.leaveGroupCallPresentation#1c50d144 returning error if any.
// Stop screen sharing in a group call
//
// Possible errors:
//
//	400 GROUPCALL_INVALID: The specified group call is invalid.
//
// See https://core.telegram.org/method/phone.leaveGroupCallPresentation for reference.
func (c *Client) PhoneLeaveGroupCallPresentation(ctx context.Context, call InputGroupCall) (UpdatesClass, error) {
	var result UpdatesBox

	request := &PhoneLeaveGroupCallPresentationRequest{
		Call: call,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
