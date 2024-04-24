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

// ReactionNotificationsFromContacts represents TL type `reactionNotificationsFromContacts#bac3a61a`.
//
// See https://core.telegram.org/constructor/reactionNotificationsFromContacts for reference.
type ReactionNotificationsFromContacts struct {
}

// ReactionNotificationsFromContactsTypeID is TL type id of ReactionNotificationsFromContacts.
const ReactionNotificationsFromContactsTypeID = 0xbac3a61a

// construct implements constructor of ReactionNotificationsFromClass.
func (r ReactionNotificationsFromContacts) construct() ReactionNotificationsFromClass { return &r }

// Ensuring interfaces in compile-time for ReactionNotificationsFromContacts.
var (
	_ bin.Encoder     = &ReactionNotificationsFromContacts{}
	_ bin.Decoder     = &ReactionNotificationsFromContacts{}
	_ bin.BareEncoder = &ReactionNotificationsFromContacts{}
	_ bin.BareDecoder = &ReactionNotificationsFromContacts{}

	_ ReactionNotificationsFromClass = &ReactionNotificationsFromContacts{}
)

func (r *ReactionNotificationsFromContacts) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReactionNotificationsFromContacts) String() string {
	if r == nil {
		return "ReactionNotificationsFromContacts(nil)"
	}
	type Alias ReactionNotificationsFromContacts
	return fmt.Sprintf("ReactionNotificationsFromContacts%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReactionNotificationsFromContacts) TypeID() uint32 {
	return ReactionNotificationsFromContactsTypeID
}

// TypeName returns name of type in TL schema.
func (*ReactionNotificationsFromContacts) TypeName() string {
	return "reactionNotificationsFromContacts"
}

// TypeInfo returns info about TL type.
func (r *ReactionNotificationsFromContacts) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reactionNotificationsFromContacts",
		ID:   ReactionNotificationsFromContactsTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReactionNotificationsFromContacts) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionNotificationsFromContacts#bac3a61a as nil")
	}
	b.PutID(ReactionNotificationsFromContactsTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReactionNotificationsFromContacts) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionNotificationsFromContacts#bac3a61a as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReactionNotificationsFromContacts) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionNotificationsFromContacts#bac3a61a to nil")
	}
	if err := b.ConsumeID(ReactionNotificationsFromContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode reactionNotificationsFromContacts#bac3a61a: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReactionNotificationsFromContacts) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionNotificationsFromContacts#bac3a61a to nil")
	}
	return nil
}

// ReactionNotificationsFromAll represents TL type `reactionNotificationsFromAll#4b9e22a0`.
//
// See https://core.telegram.org/constructor/reactionNotificationsFromAll for reference.
type ReactionNotificationsFromAll struct {
}

// ReactionNotificationsFromAllTypeID is TL type id of ReactionNotificationsFromAll.
const ReactionNotificationsFromAllTypeID = 0x4b9e22a0

// construct implements constructor of ReactionNotificationsFromClass.
func (r ReactionNotificationsFromAll) construct() ReactionNotificationsFromClass { return &r }

// Ensuring interfaces in compile-time for ReactionNotificationsFromAll.
var (
	_ bin.Encoder     = &ReactionNotificationsFromAll{}
	_ bin.Decoder     = &ReactionNotificationsFromAll{}
	_ bin.BareEncoder = &ReactionNotificationsFromAll{}
	_ bin.BareDecoder = &ReactionNotificationsFromAll{}

	_ ReactionNotificationsFromClass = &ReactionNotificationsFromAll{}
)

func (r *ReactionNotificationsFromAll) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReactionNotificationsFromAll) String() string {
	if r == nil {
		return "ReactionNotificationsFromAll(nil)"
	}
	type Alias ReactionNotificationsFromAll
	return fmt.Sprintf("ReactionNotificationsFromAll%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReactionNotificationsFromAll) TypeID() uint32 {
	return ReactionNotificationsFromAllTypeID
}

// TypeName returns name of type in TL schema.
func (*ReactionNotificationsFromAll) TypeName() string {
	return "reactionNotificationsFromAll"
}

// TypeInfo returns info about TL type.
func (r *ReactionNotificationsFromAll) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reactionNotificationsFromAll",
		ID:   ReactionNotificationsFromAllTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReactionNotificationsFromAll) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionNotificationsFromAll#4b9e22a0 as nil")
	}
	b.PutID(ReactionNotificationsFromAllTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReactionNotificationsFromAll) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionNotificationsFromAll#4b9e22a0 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReactionNotificationsFromAll) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionNotificationsFromAll#4b9e22a0 to nil")
	}
	if err := b.ConsumeID(ReactionNotificationsFromAllTypeID); err != nil {
		return fmt.Errorf("unable to decode reactionNotificationsFromAll#4b9e22a0: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReactionNotificationsFromAll) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionNotificationsFromAll#4b9e22a0 to nil")
	}
	return nil
}

// ReactionNotificationsFromClassName is schema name of ReactionNotificationsFromClass.
const ReactionNotificationsFromClassName = "ReactionNotificationsFrom"

// ReactionNotificationsFromClass represents ReactionNotificationsFrom generic type.
//
// See https://core.telegram.org/type/ReactionNotificationsFrom for reference.
//
// Example:
//
//	g, err := tg.DecodeReactionNotificationsFrom(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.ReactionNotificationsFromContacts: // reactionNotificationsFromContacts#bac3a61a
//	case *tg.ReactionNotificationsFromAll: // reactionNotificationsFromAll#4b9e22a0
//	default: panic(v)
//	}
type ReactionNotificationsFromClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ReactionNotificationsFromClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeReactionNotificationsFrom implements binary de-serialization for ReactionNotificationsFromClass.
func DecodeReactionNotificationsFrom(buf *bin.Buffer) (ReactionNotificationsFromClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ReactionNotificationsFromContactsTypeID:
		// Decoding reactionNotificationsFromContacts#bac3a61a.
		v := ReactionNotificationsFromContacts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReactionNotificationsFromClass: %w", err)
		}
		return &v, nil
	case ReactionNotificationsFromAllTypeID:
		// Decoding reactionNotificationsFromAll#4b9e22a0.
		v := ReactionNotificationsFromAll{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReactionNotificationsFromClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ReactionNotificationsFromClass: %w", bin.NewUnexpectedID(id))
	}
}

// ReactionNotificationsFrom boxes the ReactionNotificationsFromClass providing a helper.
type ReactionNotificationsFromBox struct {
	ReactionNotificationsFrom ReactionNotificationsFromClass
}

// Decode implements bin.Decoder for ReactionNotificationsFromBox.
func (b *ReactionNotificationsFromBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ReactionNotificationsFromBox to nil")
	}
	v, err := DecodeReactionNotificationsFrom(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ReactionNotificationsFrom = v
	return nil
}

// Encode implements bin.Encode for ReactionNotificationsFromBox.
func (b *ReactionNotificationsFromBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ReactionNotificationsFrom == nil {
		return fmt.Errorf("unable to encode ReactionNotificationsFromClass as nil")
	}
	return b.ReactionNotificationsFrom.Encode(buf)
}
