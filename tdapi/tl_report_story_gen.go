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

// ReportStoryRequest represents TL type `reportStory#9b5f709b`.
type ReportStoryRequest struct {
	// The identifier of the sender of the story to report
	StorySenderChatID int64
	// The identifier of the story to report
	StoryID int32
	// The reason for reporting the story
	Reason ReportReasonClass
	// Additional report details; 0-1024 characters
	Text string
}

// ReportStoryRequestTypeID is TL type id of ReportStoryRequest.
const ReportStoryRequestTypeID = 0x9b5f709b

// Ensuring interfaces in compile-time for ReportStoryRequest.
var (
	_ bin.Encoder     = &ReportStoryRequest{}
	_ bin.Decoder     = &ReportStoryRequest{}
	_ bin.BareEncoder = &ReportStoryRequest{}
	_ bin.BareDecoder = &ReportStoryRequest{}
)

func (r *ReportStoryRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.StorySenderChatID == 0) {
		return false
	}
	if !(r.StoryID == 0) {
		return false
	}
	if !(r.Reason == nil) {
		return false
	}
	if !(r.Text == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReportStoryRequest) String() string {
	if r == nil {
		return "ReportStoryRequest(nil)"
	}
	type Alias ReportStoryRequest
	return fmt.Sprintf("ReportStoryRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReportStoryRequest) TypeID() uint32 {
	return ReportStoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReportStoryRequest) TypeName() string {
	return "reportStory"
}

// TypeInfo returns info about TL type.
func (r *ReportStoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reportStory",
		ID:   ReportStoryRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "StorySenderChatID",
			SchemaName: "story_sender_chat_id",
		},
		{
			Name:       "StoryID",
			SchemaName: "story_id",
		},
		{
			Name:       "Reason",
			SchemaName: "reason",
		},
		{
			Name:       "Text",
			SchemaName: "text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReportStoryRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportStory#9b5f709b as nil")
	}
	b.PutID(ReportStoryRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReportStoryRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportStory#9b5f709b as nil")
	}
	b.PutInt53(r.StorySenderChatID)
	b.PutInt32(r.StoryID)
	if r.Reason == nil {
		return fmt.Errorf("unable to encode reportStory#9b5f709b: field reason is nil")
	}
	if err := r.Reason.Encode(b); err != nil {
		return fmt.Errorf("unable to encode reportStory#9b5f709b: field reason: %w", err)
	}
	b.PutString(r.Text)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReportStoryRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportStory#9b5f709b to nil")
	}
	if err := b.ConsumeID(ReportStoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode reportStory#9b5f709b: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReportStoryRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportStory#9b5f709b to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode reportStory#9b5f709b: field story_sender_chat_id: %w", err)
		}
		r.StorySenderChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode reportStory#9b5f709b: field story_id: %w", err)
		}
		r.StoryID = value
	}
	{
		value, err := DecodeReportReason(b)
		if err != nil {
			return fmt.Errorf("unable to decode reportStory#9b5f709b: field reason: %w", err)
		}
		r.Reason = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode reportStory#9b5f709b: field text: %w", err)
		}
		r.Text = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReportStoryRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reportStory#9b5f709b as nil")
	}
	b.ObjStart()
	b.PutID("reportStory")
	b.Comma()
	b.FieldStart("story_sender_chat_id")
	b.PutInt53(r.StorySenderChatID)
	b.Comma()
	b.FieldStart("story_id")
	b.PutInt32(r.StoryID)
	b.Comma()
	b.FieldStart("reason")
	if r.Reason == nil {
		return fmt.Errorf("unable to encode reportStory#9b5f709b: field reason is nil")
	}
	if err := r.Reason.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode reportStory#9b5f709b: field reason: %w", err)
	}
	b.Comma()
	b.FieldStart("text")
	b.PutString(r.Text)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReportStoryRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reportStory#9b5f709b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reportStory"); err != nil {
				return fmt.Errorf("unable to decode reportStory#9b5f709b: %w", err)
			}
		case "story_sender_chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode reportStory#9b5f709b: field story_sender_chat_id: %w", err)
			}
			r.StorySenderChatID = value
		case "story_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode reportStory#9b5f709b: field story_id: %w", err)
			}
			r.StoryID = value
		case "reason":
			value, err := DecodeTDLibJSONReportReason(b)
			if err != nil {
				return fmt.Errorf("unable to decode reportStory#9b5f709b: field reason: %w", err)
			}
			r.Reason = value
		case "text":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode reportStory#9b5f709b: field text: %w", err)
			}
			r.Text = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetStorySenderChatID returns value of StorySenderChatID field.
func (r *ReportStoryRequest) GetStorySenderChatID() (value int64) {
	if r == nil {
		return
	}
	return r.StorySenderChatID
}

// GetStoryID returns value of StoryID field.
func (r *ReportStoryRequest) GetStoryID() (value int32) {
	if r == nil {
		return
	}
	return r.StoryID
}

// GetReason returns value of Reason field.
func (r *ReportStoryRequest) GetReason() (value ReportReasonClass) {
	if r == nil {
		return
	}
	return r.Reason
}

// GetText returns value of Text field.
func (r *ReportStoryRequest) GetText() (value string) {
	if r == nil {
		return
	}
	return r.Text
}

// ReportStory invokes method reportStory#9b5f709b returning error if any.
func (c *Client) ReportStory(ctx context.Context, request *ReportStoryRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
