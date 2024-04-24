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

// ReactionsNotifySettings represents TL type `reactionsNotifySettings#56e34970`.
//
// See https://core.telegram.org/constructor/reactionsNotifySettings for reference.
type ReactionsNotifySettings struct {
	// Flags field of ReactionsNotifySettings.
	Flags bin.Fields
	// MessagesNotifyFrom field of ReactionsNotifySettings.
	//
	// Use SetMessagesNotifyFrom and GetMessagesNotifyFrom helpers.
	MessagesNotifyFrom ReactionNotificationsFromClass
	// StoriesNotifyFrom field of ReactionsNotifySettings.
	//
	// Use SetStoriesNotifyFrom and GetStoriesNotifyFrom helpers.
	StoriesNotifyFrom ReactionNotificationsFromClass
	// Sound field of ReactionsNotifySettings.
	Sound NotificationSoundClass
	// ShowPreviews field of ReactionsNotifySettings.
	ShowPreviews bool
}

// ReactionsNotifySettingsTypeID is TL type id of ReactionsNotifySettings.
const ReactionsNotifySettingsTypeID = 0x56e34970

// Ensuring interfaces in compile-time for ReactionsNotifySettings.
var (
	_ bin.Encoder     = &ReactionsNotifySettings{}
	_ bin.Decoder     = &ReactionsNotifySettings{}
	_ bin.BareEncoder = &ReactionsNotifySettings{}
	_ bin.BareDecoder = &ReactionsNotifySettings{}
)

func (r *ReactionsNotifySettings) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.MessagesNotifyFrom == nil) {
		return false
	}
	if !(r.StoriesNotifyFrom == nil) {
		return false
	}
	if !(r.Sound == nil) {
		return false
	}
	if !(r.ShowPreviews == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReactionsNotifySettings) String() string {
	if r == nil {
		return "ReactionsNotifySettings(nil)"
	}
	type Alias ReactionsNotifySettings
	return fmt.Sprintf("ReactionsNotifySettings%+v", Alias(*r))
}

// FillFrom fills ReactionsNotifySettings from given interface.
func (r *ReactionsNotifySettings) FillFrom(from interface {
	GetMessagesNotifyFrom() (value ReactionNotificationsFromClass, ok bool)
	GetStoriesNotifyFrom() (value ReactionNotificationsFromClass, ok bool)
	GetSound() (value NotificationSoundClass)
	GetShowPreviews() (value bool)
}) {
	if val, ok := from.GetMessagesNotifyFrom(); ok {
		r.MessagesNotifyFrom = val
	}

	if val, ok := from.GetStoriesNotifyFrom(); ok {
		r.StoriesNotifyFrom = val
	}

	r.Sound = from.GetSound()
	r.ShowPreviews = from.GetShowPreviews()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReactionsNotifySettings) TypeID() uint32 {
	return ReactionsNotifySettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*ReactionsNotifySettings) TypeName() string {
	return "reactionsNotifySettings"
}

// TypeInfo returns info about TL type.
func (r *ReactionsNotifySettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reactionsNotifySettings",
		ID:   ReactionsNotifySettingsTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MessagesNotifyFrom",
			SchemaName: "messages_notify_from",
			Null:       !r.Flags.Has(0),
		},
		{
			Name:       "StoriesNotifyFrom",
			SchemaName: "stories_notify_from",
			Null:       !r.Flags.Has(1),
		},
		{
			Name:       "Sound",
			SchemaName: "sound",
		},
		{
			Name:       "ShowPreviews",
			SchemaName: "show_previews",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *ReactionsNotifySettings) SetFlags() {
	if !(r.MessagesNotifyFrom == nil) {
		r.Flags.Set(0)
	}
	if !(r.StoriesNotifyFrom == nil) {
		r.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (r *ReactionsNotifySettings) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionsNotifySettings#56e34970 as nil")
	}
	b.PutID(ReactionsNotifySettingsTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReactionsNotifySettings) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionsNotifySettings#56e34970 as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode reactionsNotifySettings#56e34970: field flags: %w", err)
	}
	if r.Flags.Has(0) {
		if r.MessagesNotifyFrom == nil {
			return fmt.Errorf("unable to encode reactionsNotifySettings#56e34970: field messages_notify_from is nil")
		}
		if err := r.MessagesNotifyFrom.Encode(b); err != nil {
			return fmt.Errorf("unable to encode reactionsNotifySettings#56e34970: field messages_notify_from: %w", err)
		}
	}
	if r.Flags.Has(1) {
		if r.StoriesNotifyFrom == nil {
			return fmt.Errorf("unable to encode reactionsNotifySettings#56e34970: field stories_notify_from is nil")
		}
		if err := r.StoriesNotifyFrom.Encode(b); err != nil {
			return fmt.Errorf("unable to encode reactionsNotifySettings#56e34970: field stories_notify_from: %w", err)
		}
	}
	if r.Sound == nil {
		return fmt.Errorf("unable to encode reactionsNotifySettings#56e34970: field sound is nil")
	}
	if err := r.Sound.Encode(b); err != nil {
		return fmt.Errorf("unable to encode reactionsNotifySettings#56e34970: field sound: %w", err)
	}
	b.PutBool(r.ShowPreviews)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReactionsNotifySettings) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionsNotifySettings#56e34970 to nil")
	}
	if err := b.ConsumeID(ReactionsNotifySettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode reactionsNotifySettings#56e34970: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReactionsNotifySettings) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionsNotifySettings#56e34970 to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode reactionsNotifySettings#56e34970: field flags: %w", err)
		}
	}
	if r.Flags.Has(0) {
		value, err := DecodeReactionNotificationsFrom(b)
		if err != nil {
			return fmt.Errorf("unable to decode reactionsNotifySettings#56e34970: field messages_notify_from: %w", err)
		}
		r.MessagesNotifyFrom = value
	}
	if r.Flags.Has(1) {
		value, err := DecodeReactionNotificationsFrom(b)
		if err != nil {
			return fmt.Errorf("unable to decode reactionsNotifySettings#56e34970: field stories_notify_from: %w", err)
		}
		r.StoriesNotifyFrom = value
	}
	{
		value, err := DecodeNotificationSound(b)
		if err != nil {
			return fmt.Errorf("unable to decode reactionsNotifySettings#56e34970: field sound: %w", err)
		}
		r.Sound = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode reactionsNotifySettings#56e34970: field show_previews: %w", err)
		}
		r.ShowPreviews = value
	}
	return nil
}

// SetMessagesNotifyFrom sets value of MessagesNotifyFrom conditional field.
func (r *ReactionsNotifySettings) SetMessagesNotifyFrom(value ReactionNotificationsFromClass) {
	r.Flags.Set(0)
	r.MessagesNotifyFrom = value
}

// GetMessagesNotifyFrom returns value of MessagesNotifyFrom conditional field and
// boolean which is true if field was set.
func (r *ReactionsNotifySettings) GetMessagesNotifyFrom() (value ReactionNotificationsFromClass, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(0) {
		return value, false
	}
	return r.MessagesNotifyFrom, true
}

// SetStoriesNotifyFrom sets value of StoriesNotifyFrom conditional field.
func (r *ReactionsNotifySettings) SetStoriesNotifyFrom(value ReactionNotificationsFromClass) {
	r.Flags.Set(1)
	r.StoriesNotifyFrom = value
}

// GetStoriesNotifyFrom returns value of StoriesNotifyFrom conditional field and
// boolean which is true if field was set.
func (r *ReactionsNotifySettings) GetStoriesNotifyFrom() (value ReactionNotificationsFromClass, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(1) {
		return value, false
	}
	return r.StoriesNotifyFrom, true
}

// GetSound returns value of Sound field.
func (r *ReactionsNotifySettings) GetSound() (value NotificationSoundClass) {
	if r == nil {
		return
	}
	return r.Sound
}

// GetShowPreviews returns value of ShowPreviews field.
func (r *ReactionsNotifySettings) GetShowPreviews() (value bool) {
	if r == nil {
		return
	}
	return r.ShowPreviews
}
