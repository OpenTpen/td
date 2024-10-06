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

// StoriesReportRequest represents TL type `stories.report#19d8eb45`.
// Report a story.
//
// See https://core.telegram.org/method/stories.report for reference.
type StoriesReportRequest struct {
	// The peer that uploaded the story.
	Peer InputPeerClass
	// IDs of the stories to report.
	ID []int
	// Option field of StoriesReportRequest.
	Option []byte
	// Comment for report moderation
	Message string
}

// StoriesReportRequestTypeID is TL type id of StoriesReportRequest.
const StoriesReportRequestTypeID = 0x19d8eb45

// Ensuring interfaces in compile-time for StoriesReportRequest.
var (
	_ bin.Encoder     = &StoriesReportRequest{}
	_ bin.Decoder     = &StoriesReportRequest{}
	_ bin.BareEncoder = &StoriesReportRequest{}
	_ bin.BareDecoder = &StoriesReportRequest{}
)

func (r *StoriesReportRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Peer == nil) {
		return false
	}
	if !(r.ID == nil) {
		return false
	}
	if !(r.Option == nil) {
		return false
	}
	if !(r.Message == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *StoriesReportRequest) String() string {
	if r == nil {
		return "StoriesReportRequest(nil)"
	}
	type Alias StoriesReportRequest
	return fmt.Sprintf("StoriesReportRequest%+v", Alias(*r))
}

// FillFrom fills StoriesReportRequest from given interface.
func (r *StoriesReportRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetID() (value []int)
	GetOption() (value []byte)
	GetMessage() (value string)
}) {
	r.Peer = from.GetPeer()
	r.ID = from.GetID()
	r.Option = from.GetOption()
	r.Message = from.GetMessage()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesReportRequest) TypeID() uint32 {
	return StoriesReportRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesReportRequest) TypeName() string {
	return "stories.report"
}

// TypeInfo returns info about TL type.
func (r *StoriesReportRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.report",
		ID:   StoriesReportRequestTypeID,
	}
	if r == nil {
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
			Name:       "Option",
			SchemaName: "option",
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *StoriesReportRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode stories.report#19d8eb45 as nil")
	}
	b.PutID(StoriesReportRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *StoriesReportRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode stories.report#19d8eb45 as nil")
	}
	if r.Peer == nil {
		return fmt.Errorf("unable to encode stories.report#19d8eb45: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.report#19d8eb45: field peer: %w", err)
	}
	b.PutVectorHeader(len(r.ID))
	for _, v := range r.ID {
		b.PutInt(v)
	}
	b.PutBytes(r.Option)
	b.PutString(r.Message)
	return nil
}

// Decode implements bin.Decoder.
func (r *StoriesReportRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode stories.report#19d8eb45 to nil")
	}
	if err := b.ConsumeID(StoriesReportRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.report#19d8eb45: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *StoriesReportRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode stories.report#19d8eb45 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode stories.report#19d8eb45: field peer: %w", err)
		}
		r.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.report#19d8eb45: field id: %w", err)
		}

		if headerLen > 0 {
			r.ID = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode stories.report#19d8eb45: field id: %w", err)
			}
			r.ID = append(r.ID, value)
		}
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode stories.report#19d8eb45: field option: %w", err)
		}
		r.Option = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stories.report#19d8eb45: field message: %w", err)
		}
		r.Message = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (r *StoriesReportRequest) GetPeer() (value InputPeerClass) {
	if r == nil {
		return
	}
	return r.Peer
}

// GetID returns value of ID field.
func (r *StoriesReportRequest) GetID() (value []int) {
	if r == nil {
		return
	}
	return r.ID
}

// GetOption returns value of Option field.
func (r *StoriesReportRequest) GetOption() (value []byte) {
	if r == nil {
		return
	}
	return r.Option
}

// GetMessage returns value of Message field.
func (r *StoriesReportRequest) GetMessage() (value string) {
	if r == nil {
		return
	}
	return r.Message
}

// StoriesReport invokes method stories.report#19d8eb45 returning error if any.
// Report a story.
//
// Possible errors:
//
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/stories.report for reference.
func (c *Client) StoriesReport(ctx context.Context, request *StoriesReportRequest) (ReportResultClass, error) {
	var result ReportResultBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.ReportResult, nil
}
