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

// MessagesGetEmojiStickersRequest represents TL type `messages.getEmojiStickers#fbfca18f`.
// Gets the list of currently installed custom emoji stickersets¹.
//
// Links:
//  1. https://core.telegram.org/api/custom-emoji
//
// See https://core.telegram.org/method/messages.getEmojiStickers for reference.
type MessagesGetEmojiStickersRequest struct {
	// Hash used for caching, for more info click here¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int64
}

// MessagesGetEmojiStickersRequestTypeID is TL type id of MessagesGetEmojiStickersRequest.
const MessagesGetEmojiStickersRequestTypeID = 0xfbfca18f

// Ensuring interfaces in compile-time for MessagesGetEmojiStickersRequest.
var (
	_ bin.Encoder     = &MessagesGetEmojiStickersRequest{}
	_ bin.Decoder     = &MessagesGetEmojiStickersRequest{}
	_ bin.BareEncoder = &MessagesGetEmojiStickersRequest{}
	_ bin.BareDecoder = &MessagesGetEmojiStickersRequest{}
)

func (g *MessagesGetEmojiStickersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetEmojiStickersRequest) String() string {
	if g == nil {
		return "MessagesGetEmojiStickersRequest(nil)"
	}
	type Alias MessagesGetEmojiStickersRequest
	return fmt.Sprintf("MessagesGetEmojiStickersRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetEmojiStickersRequest from given interface.
func (g *MessagesGetEmojiStickersRequest) FillFrom(from interface {
	GetHash() (value int64)
}) {
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetEmojiStickersRequest) TypeID() uint32 {
	return MessagesGetEmojiStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetEmojiStickersRequest) TypeName() string {
	return "messages.getEmojiStickers"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetEmojiStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getEmojiStickers",
		ID:   MessagesGetEmojiStickersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetEmojiStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getEmojiStickers#fbfca18f as nil")
	}
	b.PutID(MessagesGetEmojiStickersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetEmojiStickersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getEmojiStickers#fbfca18f as nil")
	}
	b.PutLong(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetEmojiStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getEmojiStickers#fbfca18f to nil")
	}
	if err := b.ConsumeID(MessagesGetEmojiStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getEmojiStickers#fbfca18f: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetEmojiStickersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getEmojiStickers#fbfca18f to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getEmojiStickers#fbfca18f: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (g *MessagesGetEmojiStickersRequest) GetHash() (value int64) {
	if g == nil {
		return
	}
	return g.Hash
}

// MessagesGetEmojiStickers invokes method messages.getEmojiStickers#fbfca18f returning error if any.
// Gets the list of currently installed custom emoji stickersets¹.
//
// Links:
//  1. https://core.telegram.org/api/custom-emoji
//
// See https://core.telegram.org/method/messages.getEmojiStickers for reference.
func (c *Client) MessagesGetEmojiStickers(ctx context.Context, hash int64) (MessagesAllStickersClass, error) {
	var result MessagesAllStickersBox

	request := &MessagesGetEmojiStickersRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.AllStickers, nil
}
