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

// MessagesGetAttachMenuBotsRequest represents TL type `messages.getAttachMenuBots#16fcc2cb`.
// Returns installed attachment menu bot web apps »¹
//
// Links:
//  1. https://core.telegram.org/bots/webapps#launching-web-apps-from-the-attachment-menu
//
// See https://core.telegram.org/method/messages.getAttachMenuBots for reference.
type MessagesGetAttachMenuBotsRequest struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int64
}

// MessagesGetAttachMenuBotsRequestTypeID is TL type id of MessagesGetAttachMenuBotsRequest.
const MessagesGetAttachMenuBotsRequestTypeID = 0x16fcc2cb

// Ensuring interfaces in compile-time for MessagesGetAttachMenuBotsRequest.
var (
	_ bin.Encoder     = &MessagesGetAttachMenuBotsRequest{}
	_ bin.Decoder     = &MessagesGetAttachMenuBotsRequest{}
	_ bin.BareEncoder = &MessagesGetAttachMenuBotsRequest{}
	_ bin.BareDecoder = &MessagesGetAttachMenuBotsRequest{}
)

func (g *MessagesGetAttachMenuBotsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetAttachMenuBotsRequest) String() string {
	if g == nil {
		return "MessagesGetAttachMenuBotsRequest(nil)"
	}
	type Alias MessagesGetAttachMenuBotsRequest
	return fmt.Sprintf("MessagesGetAttachMenuBotsRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetAttachMenuBotsRequest from given interface.
func (g *MessagesGetAttachMenuBotsRequest) FillFrom(from interface {
	GetHash() (value int64)
}) {
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetAttachMenuBotsRequest) TypeID() uint32 {
	return MessagesGetAttachMenuBotsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetAttachMenuBotsRequest) TypeName() string {
	return "messages.getAttachMenuBots"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetAttachMenuBotsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getAttachMenuBots",
		ID:   MessagesGetAttachMenuBotsRequestTypeID,
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
func (g *MessagesGetAttachMenuBotsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getAttachMenuBots#16fcc2cb as nil")
	}
	b.PutID(MessagesGetAttachMenuBotsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetAttachMenuBotsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getAttachMenuBots#16fcc2cb as nil")
	}
	b.PutLong(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetAttachMenuBotsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getAttachMenuBots#16fcc2cb to nil")
	}
	if err := b.ConsumeID(MessagesGetAttachMenuBotsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getAttachMenuBots#16fcc2cb: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetAttachMenuBotsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getAttachMenuBots#16fcc2cb to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getAttachMenuBots#16fcc2cb: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (g *MessagesGetAttachMenuBotsRequest) GetHash() (value int64) {
	if g == nil {
		return
	}
	return g.Hash
}

// MessagesGetAttachMenuBots invokes method messages.getAttachMenuBots#16fcc2cb returning error if any.
// Returns installed attachment menu bot web apps »¹
//
// Links:
//  1. https://core.telegram.org/bots/webapps#launching-web-apps-from-the-attachment-menu
//
// See https://core.telegram.org/method/messages.getAttachMenuBots for reference.
func (c *Client) MessagesGetAttachMenuBots(ctx context.Context, hash int64) (AttachMenuBotsClass, error) {
	var result AttachMenuBotsBox

	request := &MessagesGetAttachMenuBotsRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.AttachMenuBots, nil
}
