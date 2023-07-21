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

// GetStoryViewersRequest represents TL type `getStoryViewers#ab175e7b`.
type GetStoryViewersRequest struct {
	// Story identifier
	StoryID int32
	// A viewer from which to return next viewers; pass null to get results from the
	// beginning
	OffsetViewer MessageViewer
	// The maximum number of story viewers to return
	Limit int32
}

// GetStoryViewersRequestTypeID is TL type id of GetStoryViewersRequest.
const GetStoryViewersRequestTypeID = 0xab175e7b

// Ensuring interfaces in compile-time for GetStoryViewersRequest.
var (
	_ bin.Encoder     = &GetStoryViewersRequest{}
	_ bin.Decoder     = &GetStoryViewersRequest{}
	_ bin.BareEncoder = &GetStoryViewersRequest{}
	_ bin.BareDecoder = &GetStoryViewersRequest{}
)

func (g *GetStoryViewersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.StoryID == 0) {
		return false
	}
	if !(g.OffsetViewer.Zero()) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetStoryViewersRequest) String() string {
	if g == nil {
		return "GetStoryViewersRequest(nil)"
	}
	type Alias GetStoryViewersRequest
	return fmt.Sprintf("GetStoryViewersRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetStoryViewersRequest) TypeID() uint32 {
	return GetStoryViewersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetStoryViewersRequest) TypeName() string {
	return "getStoryViewers"
}

// TypeInfo returns info about TL type.
func (g *GetStoryViewersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getStoryViewers",
		ID:   GetStoryViewersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "StoryID",
			SchemaName: "story_id",
		},
		{
			Name:       "OffsetViewer",
			SchemaName: "offset_viewer",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetStoryViewersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getStoryViewers#ab175e7b as nil")
	}
	b.PutID(GetStoryViewersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetStoryViewersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getStoryViewers#ab175e7b as nil")
	}
	b.PutInt32(g.StoryID)
	if err := g.OffsetViewer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getStoryViewers#ab175e7b: field offset_viewer: %w", err)
	}
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetStoryViewersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getStoryViewers#ab175e7b to nil")
	}
	if err := b.ConsumeID(GetStoryViewersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getStoryViewers#ab175e7b: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetStoryViewersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getStoryViewers#ab175e7b to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getStoryViewers#ab175e7b: field story_id: %w", err)
		}
		g.StoryID = value
	}
	{
		if err := g.OffsetViewer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode getStoryViewers#ab175e7b: field offset_viewer: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getStoryViewers#ab175e7b: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetStoryViewersRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getStoryViewers#ab175e7b as nil")
	}
	b.ObjStart()
	b.PutID("getStoryViewers")
	b.Comma()
	b.FieldStart("story_id")
	b.PutInt32(g.StoryID)
	b.Comma()
	b.FieldStart("offset_viewer")
	if err := g.OffsetViewer.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getStoryViewers#ab175e7b: field offset_viewer: %w", err)
	}
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetStoryViewersRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getStoryViewers#ab175e7b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getStoryViewers"); err != nil {
				return fmt.Errorf("unable to decode getStoryViewers#ab175e7b: %w", err)
			}
		case "story_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getStoryViewers#ab175e7b: field story_id: %w", err)
			}
			g.StoryID = value
		case "offset_viewer":
			if err := g.OffsetViewer.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode getStoryViewers#ab175e7b: field offset_viewer: %w", err)
			}
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getStoryViewers#ab175e7b: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetStoryID returns value of StoryID field.
func (g *GetStoryViewersRequest) GetStoryID() (value int32) {
	if g == nil {
		return
	}
	return g.StoryID
}

// GetOffsetViewer returns value of OffsetViewer field.
func (g *GetStoryViewersRequest) GetOffsetViewer() (value MessageViewer) {
	if g == nil {
		return
	}
	return g.OffsetViewer
}

// GetLimit returns value of Limit field.
func (g *GetStoryViewersRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetStoryViewers invokes method getStoryViewers#ab175e7b returning error if any.
func (c *Client) GetStoryViewers(ctx context.Context, request *GetStoryViewersRequest) (*MessageViewers, error) {
	var result MessageViewers

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
