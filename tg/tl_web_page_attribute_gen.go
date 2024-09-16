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

// WebPageAttributeTheme represents TL type `webPageAttributeTheme#54b56617`.
// Page theme
//
// See https://core.telegram.org/constructor/webPageAttributeTheme for reference.
type WebPageAttributeTheme struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Theme files
	//
	// Use SetDocuments and GetDocuments helpers.
	Documents []DocumentClass
	// Theme settings
	//
	// Use SetSettings and GetSettings helpers.
	Settings ThemeSettings
}

// WebPageAttributeThemeTypeID is TL type id of WebPageAttributeTheme.
const WebPageAttributeThemeTypeID = 0x54b56617

// construct implements constructor of WebPageAttributeClass.
func (w WebPageAttributeTheme) construct() WebPageAttributeClass { return &w }

// Ensuring interfaces in compile-time for WebPageAttributeTheme.
var (
	_ bin.Encoder     = &WebPageAttributeTheme{}
	_ bin.Decoder     = &WebPageAttributeTheme{}
	_ bin.BareEncoder = &WebPageAttributeTheme{}
	_ bin.BareDecoder = &WebPageAttributeTheme{}

	_ WebPageAttributeClass = &WebPageAttributeTheme{}
)

func (w *WebPageAttributeTheme) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.Flags.Zero()) {
		return false
	}
	if !(w.Documents == nil) {
		return false
	}
	if !(w.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WebPageAttributeTheme) String() string {
	if w == nil {
		return "WebPageAttributeTheme(nil)"
	}
	type Alias WebPageAttributeTheme
	return fmt.Sprintf("WebPageAttributeTheme%+v", Alias(*w))
}

// FillFrom fills WebPageAttributeTheme from given interface.
func (w *WebPageAttributeTheme) FillFrom(from interface {
	GetDocuments() (value []DocumentClass, ok bool)
	GetSettings() (value ThemeSettings, ok bool)
}) {
	if val, ok := from.GetDocuments(); ok {
		w.Documents = val
	}

	if val, ok := from.GetSettings(); ok {
		w.Settings = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WebPageAttributeTheme) TypeID() uint32 {
	return WebPageAttributeThemeTypeID
}

// TypeName returns name of type in TL schema.
func (*WebPageAttributeTheme) TypeName() string {
	return "webPageAttributeTheme"
}

// TypeInfo returns info about TL type.
func (w *WebPageAttributeTheme) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "webPageAttributeTheme",
		ID:   WebPageAttributeThemeTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Documents",
			SchemaName: "documents",
			Null:       !w.Flags.Has(0),
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
			Null:       !w.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (w *WebPageAttributeTheme) SetFlags() {
	if !(w.Documents == nil) {
		w.Flags.Set(0)
	}
	if !(w.Settings.Zero()) {
		w.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (w *WebPageAttributeTheme) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPageAttributeTheme#54b56617 as nil")
	}
	b.PutID(WebPageAttributeThemeTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *WebPageAttributeTheme) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPageAttributeTheme#54b56617 as nil")
	}
	w.SetFlags()
	if err := w.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPageAttributeTheme#54b56617: field flags: %w", err)
	}
	if w.Flags.Has(0) {
		b.PutVectorHeader(len(w.Documents))
		for idx, v := range w.Documents {
			if v == nil {
				return fmt.Errorf("unable to encode webPageAttributeTheme#54b56617: field documents element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode webPageAttributeTheme#54b56617: field documents element with index %d: %w", idx, err)
			}
		}
	}
	if w.Flags.Has(1) {
		if err := w.Settings.Encode(b); err != nil {
			return fmt.Errorf("unable to encode webPageAttributeTheme#54b56617: field settings: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (w *WebPageAttributeTheme) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPageAttributeTheme#54b56617 to nil")
	}
	if err := b.ConsumeID(WebPageAttributeThemeTypeID); err != nil {
		return fmt.Errorf("unable to decode webPageAttributeTheme#54b56617: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *WebPageAttributeTheme) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPageAttributeTheme#54b56617 to nil")
	}
	{
		if err := w.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPageAttributeTheme#54b56617: field flags: %w", err)
		}
	}
	if w.Flags.Has(0) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode webPageAttributeTheme#54b56617: field documents: %w", err)
		}

		if headerLen > 0 {
			w.Documents = make([]DocumentClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocument(b)
			if err != nil {
				return fmt.Errorf("unable to decode webPageAttributeTheme#54b56617: field documents: %w", err)
			}
			w.Documents = append(w.Documents, value)
		}
	}
	if w.Flags.Has(1) {
		if err := w.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPageAttributeTheme#54b56617: field settings: %w", err)
		}
	}
	return nil
}

// SetDocuments sets value of Documents conditional field.
func (w *WebPageAttributeTheme) SetDocuments(value []DocumentClass) {
	w.Flags.Set(0)
	w.Documents = value
}

// GetDocuments returns value of Documents conditional field and
// boolean which is true if field was set.
func (w *WebPageAttributeTheme) GetDocuments() (value []DocumentClass, ok bool) {
	if w == nil {
		return
	}
	if !w.Flags.Has(0) {
		return value, false
	}
	return w.Documents, true
}

// SetSettings sets value of Settings conditional field.
func (w *WebPageAttributeTheme) SetSettings(value ThemeSettings) {
	w.Flags.Set(1)
	w.Settings = value
}

// GetSettings returns value of Settings conditional field and
// boolean which is true if field was set.
func (w *WebPageAttributeTheme) GetSettings() (value ThemeSettings, ok bool) {
	if w == nil {
		return
	}
	if !w.Flags.Has(1) {
		return value, false
	}
	return w.Settings, true
}

// MapDocuments returns field Documents wrapped in DocumentClassArray helper.
func (w *WebPageAttributeTheme) MapDocuments() (value DocumentClassArray, ok bool) {
	if !w.Flags.Has(0) {
		return value, false
	}
	return DocumentClassArray(w.Documents), true
}

// WebPageAttributeStory represents TL type `webPageAttributeStory#2e94c3e7`.
// Webpage preview of a Telegram story
//
// See https://core.telegram.org/constructor/webPageAttributeStory for reference.
type WebPageAttributeStory struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Peer that posted the story
	Peer PeerClass
	// Story ID¹
	//
	// Links:
	//  1) https://core.telegram.org/api/stories#watching-stories
	ID int
	// May contain the story, if not the story should be fetched when and if needed using
	// stories.getStoriesByID¹ with the above id and peer.
	//
	// Links:
	//  1) https://core.telegram.org/method/stories.getStoriesByID
	//
	// Use SetStory and GetStory helpers.
	Story StoryItemClass
}

// WebPageAttributeStoryTypeID is TL type id of WebPageAttributeStory.
const WebPageAttributeStoryTypeID = 0x2e94c3e7

// construct implements constructor of WebPageAttributeClass.
func (w WebPageAttributeStory) construct() WebPageAttributeClass { return &w }

// Ensuring interfaces in compile-time for WebPageAttributeStory.
var (
	_ bin.Encoder     = &WebPageAttributeStory{}
	_ bin.Decoder     = &WebPageAttributeStory{}
	_ bin.BareEncoder = &WebPageAttributeStory{}
	_ bin.BareDecoder = &WebPageAttributeStory{}

	_ WebPageAttributeClass = &WebPageAttributeStory{}
)

func (w *WebPageAttributeStory) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.Flags.Zero()) {
		return false
	}
	if !(w.Peer == nil) {
		return false
	}
	if !(w.ID == 0) {
		return false
	}
	if !(w.Story == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WebPageAttributeStory) String() string {
	if w == nil {
		return "WebPageAttributeStory(nil)"
	}
	type Alias WebPageAttributeStory
	return fmt.Sprintf("WebPageAttributeStory%+v", Alias(*w))
}

// FillFrom fills WebPageAttributeStory from given interface.
func (w *WebPageAttributeStory) FillFrom(from interface {
	GetPeer() (value PeerClass)
	GetID() (value int)
	GetStory() (value StoryItemClass, ok bool)
}) {
	w.Peer = from.GetPeer()
	w.ID = from.GetID()
	if val, ok := from.GetStory(); ok {
		w.Story = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WebPageAttributeStory) TypeID() uint32 {
	return WebPageAttributeStoryTypeID
}

// TypeName returns name of type in TL schema.
func (*WebPageAttributeStory) TypeName() string {
	return "webPageAttributeStory"
}

// TypeInfo returns info about TL type.
func (w *WebPageAttributeStory) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "webPageAttributeStory",
		ID:   WebPageAttributeStoryTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Story",
			SchemaName: "story",
			Null:       !w.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (w *WebPageAttributeStory) SetFlags() {
	if !(w.Story == nil) {
		w.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (w *WebPageAttributeStory) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPageAttributeStory#2e94c3e7 as nil")
	}
	b.PutID(WebPageAttributeStoryTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *WebPageAttributeStory) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPageAttributeStory#2e94c3e7 as nil")
	}
	w.SetFlags()
	if err := w.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPageAttributeStory#2e94c3e7: field flags: %w", err)
	}
	if w.Peer == nil {
		return fmt.Errorf("unable to encode webPageAttributeStory#2e94c3e7: field peer is nil")
	}
	if err := w.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPageAttributeStory#2e94c3e7: field peer: %w", err)
	}
	b.PutInt(w.ID)
	if w.Flags.Has(0) {
		if w.Story == nil {
			return fmt.Errorf("unable to encode webPageAttributeStory#2e94c3e7: field story is nil")
		}
		if err := w.Story.Encode(b); err != nil {
			return fmt.Errorf("unable to encode webPageAttributeStory#2e94c3e7: field story: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (w *WebPageAttributeStory) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPageAttributeStory#2e94c3e7 to nil")
	}
	if err := b.ConsumeID(WebPageAttributeStoryTypeID); err != nil {
		return fmt.Errorf("unable to decode webPageAttributeStory#2e94c3e7: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *WebPageAttributeStory) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPageAttributeStory#2e94c3e7 to nil")
	}
	{
		if err := w.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPageAttributeStory#2e94c3e7: field flags: %w", err)
		}
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode webPageAttributeStory#2e94c3e7: field peer: %w", err)
		}
		w.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webPageAttributeStory#2e94c3e7: field id: %w", err)
		}
		w.ID = value
	}
	if w.Flags.Has(0) {
		value, err := DecodeStoryItem(b)
		if err != nil {
			return fmt.Errorf("unable to decode webPageAttributeStory#2e94c3e7: field story: %w", err)
		}
		w.Story = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (w *WebPageAttributeStory) GetPeer() (value PeerClass) {
	if w == nil {
		return
	}
	return w.Peer
}

// GetID returns value of ID field.
func (w *WebPageAttributeStory) GetID() (value int) {
	if w == nil {
		return
	}
	return w.ID
}

// SetStory sets value of Story conditional field.
func (w *WebPageAttributeStory) SetStory(value StoryItemClass) {
	w.Flags.Set(0)
	w.Story = value
}

// GetStory returns value of Story conditional field and
// boolean which is true if field was set.
func (w *WebPageAttributeStory) GetStory() (value StoryItemClass, ok bool) {
	if w == nil {
		return
	}
	if !w.Flags.Has(0) {
		return value, false
	}
	return w.Story, true
}

// WebPageAttributeStickerSet represents TL type `webPageAttributeStickerSet#50cc03d3`.
// Contains info about a stickerset »¹, for a webPage² preview of a stickerset deep
// link »³ (the webPage⁴ will have a type of telegram_stickerset).
//
// Links:
//  1. https://core.telegram.org/api/stickers
//  2. https://core.telegram.org/constructor/webPage
//  3. https://core.telegram.org/api/links#stickerset-links
//  4. https://core.telegram.org/constructor/webPage
//
// See https://core.telegram.org/constructor/webPageAttributeStickerSet for reference.
type WebPageAttributeStickerSet struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this i s a custom emoji stickerset¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/custom-emoji
	Emojis bool
	// Whether the color of this TGS custom emoji stickerset should be changed to the text
	// color when used in messages, the accent color if used as emoji status, white on chat
	// photos, or another appropriate color based on context.
	TextColor bool
	// A subset of the stickerset in the stickerset.
	Stickers []DocumentClass
}

// WebPageAttributeStickerSetTypeID is TL type id of WebPageAttributeStickerSet.
const WebPageAttributeStickerSetTypeID = 0x50cc03d3

// construct implements constructor of WebPageAttributeClass.
func (w WebPageAttributeStickerSet) construct() WebPageAttributeClass { return &w }

// Ensuring interfaces in compile-time for WebPageAttributeStickerSet.
var (
	_ bin.Encoder     = &WebPageAttributeStickerSet{}
	_ bin.Decoder     = &WebPageAttributeStickerSet{}
	_ bin.BareEncoder = &WebPageAttributeStickerSet{}
	_ bin.BareDecoder = &WebPageAttributeStickerSet{}

	_ WebPageAttributeClass = &WebPageAttributeStickerSet{}
)

func (w *WebPageAttributeStickerSet) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.Flags.Zero()) {
		return false
	}
	if !(w.Emojis == false) {
		return false
	}
	if !(w.TextColor == false) {
		return false
	}
	if !(w.Stickers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WebPageAttributeStickerSet) String() string {
	if w == nil {
		return "WebPageAttributeStickerSet(nil)"
	}
	type Alias WebPageAttributeStickerSet
	return fmt.Sprintf("WebPageAttributeStickerSet%+v", Alias(*w))
}

// FillFrom fills WebPageAttributeStickerSet from given interface.
func (w *WebPageAttributeStickerSet) FillFrom(from interface {
	GetEmojis() (value bool)
	GetTextColor() (value bool)
	GetStickers() (value []DocumentClass)
}) {
	w.Emojis = from.GetEmojis()
	w.TextColor = from.GetTextColor()
	w.Stickers = from.GetStickers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WebPageAttributeStickerSet) TypeID() uint32 {
	return WebPageAttributeStickerSetTypeID
}

// TypeName returns name of type in TL schema.
func (*WebPageAttributeStickerSet) TypeName() string {
	return "webPageAttributeStickerSet"
}

// TypeInfo returns info about TL type.
func (w *WebPageAttributeStickerSet) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "webPageAttributeStickerSet",
		ID:   WebPageAttributeStickerSetTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Emojis",
			SchemaName: "emojis",
			Null:       !w.Flags.Has(0),
		},
		{
			Name:       "TextColor",
			SchemaName: "text_color",
			Null:       !w.Flags.Has(1),
		},
		{
			Name:       "Stickers",
			SchemaName: "stickers",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (w *WebPageAttributeStickerSet) SetFlags() {
	if !(w.Emojis == false) {
		w.Flags.Set(0)
	}
	if !(w.TextColor == false) {
		w.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (w *WebPageAttributeStickerSet) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPageAttributeStickerSet#50cc03d3 as nil")
	}
	b.PutID(WebPageAttributeStickerSetTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *WebPageAttributeStickerSet) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPageAttributeStickerSet#50cc03d3 as nil")
	}
	w.SetFlags()
	if err := w.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPageAttributeStickerSet#50cc03d3: field flags: %w", err)
	}
	b.PutVectorHeader(len(w.Stickers))
	for idx, v := range w.Stickers {
		if v == nil {
			return fmt.Errorf("unable to encode webPageAttributeStickerSet#50cc03d3: field stickers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode webPageAttributeStickerSet#50cc03d3: field stickers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (w *WebPageAttributeStickerSet) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPageAttributeStickerSet#50cc03d3 to nil")
	}
	if err := b.ConsumeID(WebPageAttributeStickerSetTypeID); err != nil {
		return fmt.Errorf("unable to decode webPageAttributeStickerSet#50cc03d3: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *WebPageAttributeStickerSet) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPageAttributeStickerSet#50cc03d3 to nil")
	}
	{
		if err := w.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPageAttributeStickerSet#50cc03d3: field flags: %w", err)
		}
	}
	w.Emojis = w.Flags.Has(0)
	w.TextColor = w.Flags.Has(1)
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode webPageAttributeStickerSet#50cc03d3: field stickers: %w", err)
		}

		if headerLen > 0 {
			w.Stickers = make([]DocumentClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocument(b)
			if err != nil {
				return fmt.Errorf("unable to decode webPageAttributeStickerSet#50cc03d3: field stickers: %w", err)
			}
			w.Stickers = append(w.Stickers, value)
		}
	}
	return nil
}

// SetEmojis sets value of Emojis conditional field.
func (w *WebPageAttributeStickerSet) SetEmojis(value bool) {
	if value {
		w.Flags.Set(0)
		w.Emojis = true
	} else {
		w.Flags.Unset(0)
		w.Emojis = false
	}
}

// GetEmojis returns value of Emojis conditional field.
func (w *WebPageAttributeStickerSet) GetEmojis() (value bool) {
	if w == nil {
		return
	}
	return w.Flags.Has(0)
}

// SetTextColor sets value of TextColor conditional field.
func (w *WebPageAttributeStickerSet) SetTextColor(value bool) {
	if value {
		w.Flags.Set(1)
		w.TextColor = true
	} else {
		w.Flags.Unset(1)
		w.TextColor = false
	}
}

// GetTextColor returns value of TextColor conditional field.
func (w *WebPageAttributeStickerSet) GetTextColor() (value bool) {
	if w == nil {
		return
	}
	return w.Flags.Has(1)
}

// GetStickers returns value of Stickers field.
func (w *WebPageAttributeStickerSet) GetStickers() (value []DocumentClass) {
	if w == nil {
		return
	}
	return w.Stickers
}

// MapStickers returns field Stickers wrapped in DocumentClassArray helper.
func (w *WebPageAttributeStickerSet) MapStickers() (value DocumentClassArray) {
	return DocumentClassArray(w.Stickers)
}

// WebPageAttributeClassName is schema name of WebPageAttributeClass.
const WebPageAttributeClassName = "WebPageAttribute"

// WebPageAttributeClass represents WebPageAttribute generic type.
//
// See https://core.telegram.org/type/WebPageAttribute for reference.
//
// Example:
//
//	g, err := tg.DecodeWebPageAttribute(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.WebPageAttributeTheme: // webPageAttributeTheme#54b56617
//	case *tg.WebPageAttributeStory: // webPageAttributeStory#2e94c3e7
//	case *tg.WebPageAttributeStickerSet: // webPageAttributeStickerSet#50cc03d3
//	default: panic(v)
//	}
type WebPageAttributeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() WebPageAttributeClass

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

// DecodeWebPageAttribute implements binary de-serialization for WebPageAttributeClass.
func DecodeWebPageAttribute(buf *bin.Buffer) (WebPageAttributeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case WebPageAttributeThemeTypeID:
		// Decoding webPageAttributeTheme#54b56617.
		v := WebPageAttributeTheme{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WebPageAttributeClass: %w", err)
		}
		return &v, nil
	case WebPageAttributeStoryTypeID:
		// Decoding webPageAttributeStory#2e94c3e7.
		v := WebPageAttributeStory{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WebPageAttributeClass: %w", err)
		}
		return &v, nil
	case WebPageAttributeStickerSetTypeID:
		// Decoding webPageAttributeStickerSet#50cc03d3.
		v := WebPageAttributeStickerSet{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WebPageAttributeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode WebPageAttributeClass: %w", bin.NewUnexpectedID(id))
	}
}

// WebPageAttribute boxes the WebPageAttributeClass providing a helper.
type WebPageAttributeBox struct {
	WebPageAttribute WebPageAttributeClass
}

// Decode implements bin.Decoder for WebPageAttributeBox.
func (b *WebPageAttributeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode WebPageAttributeBox to nil")
	}
	v, err := DecodeWebPageAttribute(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.WebPageAttribute = v
	return nil
}

// Encode implements bin.Encode for WebPageAttributeBox.
func (b *WebPageAttributeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.WebPageAttribute == nil {
		return fmt.Errorf("unable to encode WebPageAttributeClass as nil")
	}
	return b.WebPageAttribute.Encode(buf)
}
