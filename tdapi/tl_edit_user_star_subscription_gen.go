// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// EditUserStarSubscriptionRequest represents TL type `editUserStarSubscription#51b16e89`.
type EditUserStarSubscriptionRequest struct {
	// User identifier
	UserID int64
	// Telegram payment identifier of the subscription
	TelegramPaymentChargeID string
	// Pass true to cancel the subscription; pass false to allow the user to enable it
	IsCanceled bool
}

// EditUserStarSubscriptionRequestTypeID is TL type id of EditUserStarSubscriptionRequest.
const EditUserStarSubscriptionRequestTypeID = 0x51b16e89

// Ensuring interfaces in compile-time for EditUserStarSubscriptionRequest.
var (
	_ bin.Encoder     = &EditUserStarSubscriptionRequest{}
	_ bin.Decoder     = &EditUserStarSubscriptionRequest{}
	_ bin.BareEncoder = &EditUserStarSubscriptionRequest{}
	_ bin.BareDecoder = &EditUserStarSubscriptionRequest{}
)

func (e *EditUserStarSubscriptionRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.UserID == 0) {
		return false
	}
	if !(e.TelegramPaymentChargeID == "") {
		return false
	}
	if !(e.IsCanceled == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EditUserStarSubscriptionRequest) String() string {
	if e == nil {
		return "EditUserStarSubscriptionRequest(nil)"
	}
	type Alias EditUserStarSubscriptionRequest
	return fmt.Sprintf("EditUserStarSubscriptionRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EditUserStarSubscriptionRequest) TypeID() uint32 {
	return EditUserStarSubscriptionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EditUserStarSubscriptionRequest) TypeName() string {
	return "editUserStarSubscription"
}

// TypeInfo returns info about TL type.
func (e *EditUserStarSubscriptionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "editUserStarSubscription",
		ID:   EditUserStarSubscriptionRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "TelegramPaymentChargeID",
			SchemaName: "telegram_payment_charge_id",
		},
		{
			Name:       "IsCanceled",
			SchemaName: "is_canceled",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EditUserStarSubscriptionRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editUserStarSubscription#51b16e89 as nil")
	}
	b.PutID(EditUserStarSubscriptionRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EditUserStarSubscriptionRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editUserStarSubscription#51b16e89 as nil")
	}
	b.PutInt53(e.UserID)
	b.PutString(e.TelegramPaymentChargeID)
	b.PutBool(e.IsCanceled)
	return nil
}

// Decode implements bin.Decoder.
func (e *EditUserStarSubscriptionRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editUserStarSubscription#51b16e89 to nil")
	}
	if err := b.ConsumeID(EditUserStarSubscriptionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode editUserStarSubscription#51b16e89: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EditUserStarSubscriptionRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editUserStarSubscription#51b16e89 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode editUserStarSubscription#51b16e89: field user_id: %w", err)
		}
		e.UserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode editUserStarSubscription#51b16e89: field telegram_payment_charge_id: %w", err)
		}
		e.TelegramPaymentChargeID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode editUserStarSubscription#51b16e89: field is_canceled: %w", err)
		}
		e.IsCanceled = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EditUserStarSubscriptionRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode editUserStarSubscription#51b16e89 as nil")
	}
	b.ObjStart()
	b.PutID("editUserStarSubscription")
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(e.UserID)
	b.Comma()
	b.FieldStart("telegram_payment_charge_id")
	b.PutString(e.TelegramPaymentChargeID)
	b.Comma()
	b.FieldStart("is_canceled")
	b.PutBool(e.IsCanceled)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EditUserStarSubscriptionRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode editUserStarSubscription#51b16e89 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("editUserStarSubscription"); err != nil {
				return fmt.Errorf("unable to decode editUserStarSubscription#51b16e89: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode editUserStarSubscription#51b16e89: field user_id: %w", err)
			}
			e.UserID = value
		case "telegram_payment_charge_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode editUserStarSubscription#51b16e89: field telegram_payment_charge_id: %w", err)
			}
			e.TelegramPaymentChargeID = value
		case "is_canceled":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode editUserStarSubscription#51b16e89: field is_canceled: %w", err)
			}
			e.IsCanceled = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (e *EditUserStarSubscriptionRequest) GetUserID() (value int64) {
	if e == nil {
		return
	}
	return e.UserID
}

// GetTelegramPaymentChargeID returns value of TelegramPaymentChargeID field.
func (e *EditUserStarSubscriptionRequest) GetTelegramPaymentChargeID() (value string) {
	if e == nil {
		return
	}
	return e.TelegramPaymentChargeID
}

// GetIsCanceled returns value of IsCanceled field.
func (e *EditUserStarSubscriptionRequest) GetIsCanceled() (value bool) {
	if e == nil {
		return
	}
	return e.IsCanceled
}

// EditUserStarSubscription invokes method editUserStarSubscription#51b16e89 returning error if any.
func (c *Client) EditUserStarSubscription(ctx context.Context, request *EditUserStarSubscriptionRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
