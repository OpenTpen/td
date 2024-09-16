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

// Boost represents TL type `boost#4b3e14d6`.
// Info about one or more boosts¹ applied by a specific user.
//
// Links:
//  1. https://core.telegram.org/api/boost
//
// See https://core.telegram.org/constructor/boost for reference.
type Boost struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this boost was applied because the channel/supergroup directly gifted a
	// subscription to the user¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/giveaways
	Gift bool
	// Whether this boost was applied because the user was chosen in a giveaway started by
	// the channel/supergroup¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/giveaways
	Giveaway bool
	// If set, the user hasn't yet invoked payments.applyGiftCode¹ to claim a subscription
	// gifted directly or in a giveaway by the channel².
	//
	// Links:
	//  1) https://core.telegram.org/method/payments.applyGiftCode
	//  2) https://core.telegram.org/api/giveaways
	Unclaimed bool
	// Unique ID for this set of boosts.
	ID string
	// ID of the user that applied the boost.
	//
	// Use SetUserID and GetUserID helpers.
	UserID int64
	// The message ID of the giveaway¹
	//
	// Links:
	//  1) https://core.telegram.org/api/giveaways
	//
	// Use SetGiveawayMsgID and GetGiveawayMsgID helpers.
	GiveawayMsgID int
	// When was the boost applied
	Date int
	// When does the boost expire
	Expires int
	// The created Telegram Premium gift code, only set if either gift or giveaway are set
	// AND it is either a gift code for the currently logged in user or if it was already
	// claimed.
	//
	// Use SetUsedGiftSlug and GetUsedGiftSlug helpers.
	UsedGiftSlug string
	// If set, this boost counts as multiplier boosts, otherwise it counts as a single boost.
	//
	// Use SetMultiplier and GetMultiplier helpers.
	Multiplier int
	// Stars field of Boost.
	//
	// Use SetStars and GetStars helpers.
	Stars int64
}

// BoostTypeID is TL type id of Boost.
const BoostTypeID = 0x4b3e14d6

// Ensuring interfaces in compile-time for Boost.
var (
	_ bin.Encoder     = &Boost{}
	_ bin.Decoder     = &Boost{}
	_ bin.BareEncoder = &Boost{}
	_ bin.BareDecoder = &Boost{}
)

func (b *Boost) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Flags.Zero()) {
		return false
	}
	if !(b.Gift == false) {
		return false
	}
	if !(b.Giveaway == false) {
		return false
	}
	if !(b.Unclaimed == false) {
		return false
	}
	if !(b.ID == "") {
		return false
	}
	if !(b.UserID == 0) {
		return false
	}
	if !(b.GiveawayMsgID == 0) {
		return false
	}
	if !(b.Date == 0) {
		return false
	}
	if !(b.Expires == 0) {
		return false
	}
	if !(b.UsedGiftSlug == "") {
		return false
	}
	if !(b.Multiplier == 0) {
		return false
	}
	if !(b.Stars == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *Boost) String() string {
	if b == nil {
		return "Boost(nil)"
	}
	type Alias Boost
	return fmt.Sprintf("Boost%+v", Alias(*b))
}

// FillFrom fills Boost from given interface.
func (b *Boost) FillFrom(from interface {
	GetGift() (value bool)
	GetGiveaway() (value bool)
	GetUnclaimed() (value bool)
	GetID() (value string)
	GetUserID() (value int64, ok bool)
	GetGiveawayMsgID() (value int, ok bool)
	GetDate() (value int)
	GetExpires() (value int)
	GetUsedGiftSlug() (value string, ok bool)
	GetMultiplier() (value int, ok bool)
	GetStars() (value int64, ok bool)
}) {
	b.Gift = from.GetGift()
	b.Giveaway = from.GetGiveaway()
	b.Unclaimed = from.GetUnclaimed()
	b.ID = from.GetID()
	if val, ok := from.GetUserID(); ok {
		b.UserID = val
	}

	if val, ok := from.GetGiveawayMsgID(); ok {
		b.GiveawayMsgID = val
	}

	b.Date = from.GetDate()
	b.Expires = from.GetExpires()
	if val, ok := from.GetUsedGiftSlug(); ok {
		b.UsedGiftSlug = val
	}

	if val, ok := from.GetMultiplier(); ok {
		b.Multiplier = val
	}

	if val, ok := from.GetStars(); ok {
		b.Stars = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Boost) TypeID() uint32 {
	return BoostTypeID
}

// TypeName returns name of type in TL schema.
func (*Boost) TypeName() string {
	return "boost"
}

// TypeInfo returns info about TL type.
func (b *Boost) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "boost",
		ID:   BoostTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Gift",
			SchemaName: "gift",
			Null:       !b.Flags.Has(1),
		},
		{
			Name:       "Giveaway",
			SchemaName: "giveaway",
			Null:       !b.Flags.Has(2),
		},
		{
			Name:       "Unclaimed",
			SchemaName: "unclaimed",
			Null:       !b.Flags.Has(3),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
			Null:       !b.Flags.Has(0),
		},
		{
			Name:       "GiveawayMsgID",
			SchemaName: "giveaway_msg_id",
			Null:       !b.Flags.Has(2),
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "Expires",
			SchemaName: "expires",
		},
		{
			Name:       "UsedGiftSlug",
			SchemaName: "used_gift_slug",
			Null:       !b.Flags.Has(4),
		},
		{
			Name:       "Multiplier",
			SchemaName: "multiplier",
			Null:       !b.Flags.Has(5),
		},
		{
			Name:       "Stars",
			SchemaName: "stars",
			Null:       !b.Flags.Has(6),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (b *Boost) SetFlags() {
	if !(b.Gift == false) {
		b.Flags.Set(1)
	}
	if !(b.Giveaway == false) {
		b.Flags.Set(2)
	}
	if !(b.Unclaimed == false) {
		b.Flags.Set(3)
	}
	if !(b.UserID == 0) {
		b.Flags.Set(0)
	}
	if !(b.GiveawayMsgID == 0) {
		b.Flags.Set(2)
	}
	if !(b.UsedGiftSlug == "") {
		b.Flags.Set(4)
	}
	if !(b.Multiplier == 0) {
		b.Flags.Set(5)
	}
	if !(b.Stars == 0) {
		b.Flags.Set(6)
	}
}

// Encode implements bin.Encoder.
func (b *Boost) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode boost#4b3e14d6 as nil")
	}
	buf.PutID(BoostTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *Boost) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode boost#4b3e14d6 as nil")
	}
	b.SetFlags()
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode boost#4b3e14d6: field flags: %w", err)
	}
	buf.PutString(b.ID)
	if b.Flags.Has(0) {
		buf.PutLong(b.UserID)
	}
	if b.Flags.Has(2) {
		buf.PutInt(b.GiveawayMsgID)
	}
	buf.PutInt(b.Date)
	buf.PutInt(b.Expires)
	if b.Flags.Has(4) {
		buf.PutString(b.UsedGiftSlug)
	}
	if b.Flags.Has(5) {
		buf.PutInt(b.Multiplier)
	}
	if b.Flags.Has(6) {
		buf.PutLong(b.Stars)
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *Boost) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode boost#4b3e14d6 to nil")
	}
	if err := buf.ConsumeID(BoostTypeID); err != nil {
		return fmt.Errorf("unable to decode boost#4b3e14d6: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *Boost) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode boost#4b3e14d6 to nil")
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode boost#4b3e14d6: field flags: %w", err)
		}
	}
	b.Gift = b.Flags.Has(1)
	b.Giveaway = b.Flags.Has(2)
	b.Unclaimed = b.Flags.Has(3)
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode boost#4b3e14d6: field id: %w", err)
		}
		b.ID = value
	}
	if b.Flags.Has(0) {
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode boost#4b3e14d6: field user_id: %w", err)
		}
		b.UserID = value
	}
	if b.Flags.Has(2) {
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode boost#4b3e14d6: field giveaway_msg_id: %w", err)
		}
		b.GiveawayMsgID = value
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode boost#4b3e14d6: field date: %w", err)
		}
		b.Date = value
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode boost#4b3e14d6: field expires: %w", err)
		}
		b.Expires = value
	}
	if b.Flags.Has(4) {
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode boost#4b3e14d6: field used_gift_slug: %w", err)
		}
		b.UsedGiftSlug = value
	}
	if b.Flags.Has(5) {
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode boost#4b3e14d6: field multiplier: %w", err)
		}
		b.Multiplier = value
	}
	if b.Flags.Has(6) {
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode boost#4b3e14d6: field stars: %w", err)
		}
		b.Stars = value
	}
	return nil
}

// SetGift sets value of Gift conditional field.
func (b *Boost) SetGift(value bool) {
	if value {
		b.Flags.Set(1)
		b.Gift = true
	} else {
		b.Flags.Unset(1)
		b.Gift = false
	}
}

// GetGift returns value of Gift conditional field.
func (b *Boost) GetGift() (value bool) {
	if b == nil {
		return
	}
	return b.Flags.Has(1)
}

// SetGiveaway sets value of Giveaway conditional field.
func (b *Boost) SetGiveaway(value bool) {
	if value {
		b.Flags.Set(2)
		b.Giveaway = true
	} else {
		b.Flags.Unset(2)
		b.Giveaway = false
	}
}

// GetGiveaway returns value of Giveaway conditional field.
func (b *Boost) GetGiveaway() (value bool) {
	if b == nil {
		return
	}
	return b.Flags.Has(2)
}

// SetUnclaimed sets value of Unclaimed conditional field.
func (b *Boost) SetUnclaimed(value bool) {
	if value {
		b.Flags.Set(3)
		b.Unclaimed = true
	} else {
		b.Flags.Unset(3)
		b.Unclaimed = false
	}
}

// GetUnclaimed returns value of Unclaimed conditional field.
func (b *Boost) GetUnclaimed() (value bool) {
	if b == nil {
		return
	}
	return b.Flags.Has(3)
}

// GetID returns value of ID field.
func (b *Boost) GetID() (value string) {
	if b == nil {
		return
	}
	return b.ID
}

// SetUserID sets value of UserID conditional field.
func (b *Boost) SetUserID(value int64) {
	b.Flags.Set(0)
	b.UserID = value
}

// GetUserID returns value of UserID conditional field and
// boolean which is true if field was set.
func (b *Boost) GetUserID() (value int64, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(0) {
		return value, false
	}
	return b.UserID, true
}

// SetGiveawayMsgID sets value of GiveawayMsgID conditional field.
func (b *Boost) SetGiveawayMsgID(value int) {
	b.Flags.Set(2)
	b.GiveawayMsgID = value
}

// GetGiveawayMsgID returns value of GiveawayMsgID conditional field and
// boolean which is true if field was set.
func (b *Boost) GetGiveawayMsgID() (value int, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(2) {
		return value, false
	}
	return b.GiveawayMsgID, true
}

// GetDate returns value of Date field.
func (b *Boost) GetDate() (value int) {
	if b == nil {
		return
	}
	return b.Date
}

// GetExpires returns value of Expires field.
func (b *Boost) GetExpires() (value int) {
	if b == nil {
		return
	}
	return b.Expires
}

// SetUsedGiftSlug sets value of UsedGiftSlug conditional field.
func (b *Boost) SetUsedGiftSlug(value string) {
	b.Flags.Set(4)
	b.UsedGiftSlug = value
}

// GetUsedGiftSlug returns value of UsedGiftSlug conditional field and
// boolean which is true if field was set.
func (b *Boost) GetUsedGiftSlug() (value string, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(4) {
		return value, false
	}
	return b.UsedGiftSlug, true
}

// SetMultiplier sets value of Multiplier conditional field.
func (b *Boost) SetMultiplier(value int) {
	b.Flags.Set(5)
	b.Multiplier = value
}

// GetMultiplier returns value of Multiplier conditional field and
// boolean which is true if field was set.
func (b *Boost) GetMultiplier() (value int, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(5) {
		return value, false
	}
	return b.Multiplier, true
}

// SetStars sets value of Stars conditional field.
func (b *Boost) SetStars(value int64) {
	b.Flags.Set(6)
	b.Stars = value
}

// GetStars returns value of Stars conditional field and
// boolean which is true if field was set.
func (b *Boost) GetStars() (value int64, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(6) {
		return value, false
	}
	return b.Stars, true
}
