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

// Poll represents TL type `poll#8c34a82a`.
type Poll struct {
	// Unique poll identifier
	ID int64
	// Poll question; 1-300 characters
	Question string
	// List of poll answer options
	Options []PollOption
	// Total number of voters, participating in the poll
	TotalVoterCount int32
	// Identifiers of recent voters, if the poll is non-anonymous
	RecentVoterIDs []MessageSenderClass
	// True, if the poll is anonymous
	IsAnonymous bool
	// Type of the poll
	Type PollTypeClass
	// Amount of time the poll will be active after creation, in seconds
	OpenPeriod int32
	// Point in time (Unix timestamp) when the poll will automatically be closed
	CloseDate int32
	// True, if the poll is closed
	IsClosed bool
}

// PollTypeID is TL type id of Poll.
const PollTypeID = 0x8c34a82a

// Ensuring interfaces in compile-time for Poll.
var (
	_ bin.Encoder     = &Poll{}
	_ bin.Decoder     = &Poll{}
	_ bin.BareEncoder = &Poll{}
	_ bin.BareDecoder = &Poll{}
)

func (p *Poll) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ID == 0) {
		return false
	}
	if !(p.Question == "") {
		return false
	}
	if !(p.Options == nil) {
		return false
	}
	if !(p.TotalVoterCount == 0) {
		return false
	}
	if !(p.RecentVoterIDs == nil) {
		return false
	}
	if !(p.IsAnonymous == false) {
		return false
	}
	if !(p.Type == nil) {
		return false
	}
	if !(p.OpenPeriod == 0) {
		return false
	}
	if !(p.CloseDate == 0) {
		return false
	}
	if !(p.IsClosed == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *Poll) String() string {
	if p == nil {
		return "Poll(nil)"
	}
	type Alias Poll
	return fmt.Sprintf("Poll%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Poll) TypeID() uint32 {
	return PollTypeID
}

// TypeName returns name of type in TL schema.
func (*Poll) TypeName() string {
	return "poll"
}

// TypeInfo returns info about TL type.
func (p *Poll) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "poll",
		ID:   PollTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Question",
			SchemaName: "question",
		},
		{
			Name:       "Options",
			SchemaName: "options",
		},
		{
			Name:       "TotalVoterCount",
			SchemaName: "total_voter_count",
		},
		{
			Name:       "RecentVoterIDs",
			SchemaName: "recent_voter_ids",
		},
		{
			Name:       "IsAnonymous",
			SchemaName: "is_anonymous",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "OpenPeriod",
			SchemaName: "open_period",
		},
		{
			Name:       "CloseDate",
			SchemaName: "close_date",
		},
		{
			Name:       "IsClosed",
			SchemaName: "is_closed",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *Poll) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode poll#8c34a82a as nil")
	}
	b.PutID(PollTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *Poll) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode poll#8c34a82a as nil")
	}
	b.PutLong(p.ID)
	b.PutString(p.Question)
	b.PutInt(len(p.Options))
	for idx, v := range p.Options {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare poll#8c34a82a: field options element with index %d: %w", idx, err)
		}
	}
	b.PutInt32(p.TotalVoterCount)
	b.PutInt(len(p.RecentVoterIDs))
	for idx, v := range p.RecentVoterIDs {
		if v == nil {
			return fmt.Errorf("unable to encode poll#8c34a82a: field recent_voter_ids element with index %d is nil", idx)
		}
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare poll#8c34a82a: field recent_voter_ids element with index %d: %w", idx, err)
		}
	}
	b.PutBool(p.IsAnonymous)
	if p.Type == nil {
		return fmt.Errorf("unable to encode poll#8c34a82a: field type is nil")
	}
	if err := p.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode poll#8c34a82a: field type: %w", err)
	}
	b.PutInt32(p.OpenPeriod)
	b.PutInt32(p.CloseDate)
	b.PutBool(p.IsClosed)
	return nil
}

// Decode implements bin.Decoder.
func (p *Poll) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode poll#8c34a82a to nil")
	}
	if err := b.ConsumeID(PollTypeID); err != nil {
		return fmt.Errorf("unable to decode poll#8c34a82a: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *Poll) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode poll#8c34a82a to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode poll#8c34a82a: field id: %w", err)
		}
		p.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode poll#8c34a82a: field question: %w", err)
		}
		p.Question = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode poll#8c34a82a: field options: %w", err)
		}

		if headerLen > 0 {
			p.Options = make([]PollOption, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PollOption
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare poll#8c34a82a: field options: %w", err)
			}
			p.Options = append(p.Options, value)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode poll#8c34a82a: field total_voter_count: %w", err)
		}
		p.TotalVoterCount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode poll#8c34a82a: field recent_voter_ids: %w", err)
		}

		if headerLen > 0 {
			p.RecentVoterIDs = make([]MessageSenderClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageSender(b)
			if err != nil {
				return fmt.Errorf("unable to decode poll#8c34a82a: field recent_voter_ids: %w", err)
			}
			p.RecentVoterIDs = append(p.RecentVoterIDs, value)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode poll#8c34a82a: field is_anonymous: %w", err)
		}
		p.IsAnonymous = value
	}
	{
		value, err := DecodePollType(b)
		if err != nil {
			return fmt.Errorf("unable to decode poll#8c34a82a: field type: %w", err)
		}
		p.Type = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode poll#8c34a82a: field open_period: %w", err)
		}
		p.OpenPeriod = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode poll#8c34a82a: field close_date: %w", err)
		}
		p.CloseDate = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode poll#8c34a82a: field is_closed: %w", err)
		}
		p.IsClosed = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *Poll) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode poll#8c34a82a as nil")
	}
	b.ObjStart()
	b.PutID("poll")
	b.Comma()
	b.FieldStart("id")
	b.PutLong(p.ID)
	b.Comma()
	b.FieldStart("question")
	b.PutString(p.Question)
	b.Comma()
	b.FieldStart("options")
	b.ArrStart()
	for idx, v := range p.Options {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode poll#8c34a82a: field options element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("total_voter_count")
	b.PutInt32(p.TotalVoterCount)
	b.Comma()
	b.FieldStart("recent_voter_ids")
	b.ArrStart()
	for idx, v := range p.RecentVoterIDs {
		if v == nil {
			return fmt.Errorf("unable to encode poll#8c34a82a: field recent_voter_ids element with index %d is nil", idx)
		}
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode poll#8c34a82a: field recent_voter_ids element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("is_anonymous")
	b.PutBool(p.IsAnonymous)
	b.Comma()
	b.FieldStart("type")
	if p.Type == nil {
		return fmt.Errorf("unable to encode poll#8c34a82a: field type is nil")
	}
	if err := p.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode poll#8c34a82a: field type: %w", err)
	}
	b.Comma()
	b.FieldStart("open_period")
	b.PutInt32(p.OpenPeriod)
	b.Comma()
	b.FieldStart("close_date")
	b.PutInt32(p.CloseDate)
	b.Comma()
	b.FieldStart("is_closed")
	b.PutBool(p.IsClosed)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *Poll) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode poll#8c34a82a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("poll"); err != nil {
				return fmt.Errorf("unable to decode poll#8c34a82a: %w", err)
			}
		case "id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode poll#8c34a82a: field id: %w", err)
			}
			p.ID = value
		case "question":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode poll#8c34a82a: field question: %w", err)
			}
			p.Question = value
		case "options":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value PollOption
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode poll#8c34a82a: field options: %w", err)
				}
				p.Options = append(p.Options, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode poll#8c34a82a: field options: %w", err)
			}
		case "total_voter_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode poll#8c34a82a: field total_voter_count: %w", err)
			}
			p.TotalVoterCount = value
		case "recent_voter_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := DecodeTDLibJSONMessageSender(b)
				if err != nil {
					return fmt.Errorf("unable to decode poll#8c34a82a: field recent_voter_ids: %w", err)
				}
				p.RecentVoterIDs = append(p.RecentVoterIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode poll#8c34a82a: field recent_voter_ids: %w", err)
			}
		case "is_anonymous":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode poll#8c34a82a: field is_anonymous: %w", err)
			}
			p.IsAnonymous = value
		case "type":
			value, err := DecodeTDLibJSONPollType(b)
			if err != nil {
				return fmt.Errorf("unable to decode poll#8c34a82a: field type: %w", err)
			}
			p.Type = value
		case "open_period":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode poll#8c34a82a: field open_period: %w", err)
			}
			p.OpenPeriod = value
		case "close_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode poll#8c34a82a: field close_date: %w", err)
			}
			p.CloseDate = value
		case "is_closed":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode poll#8c34a82a: field is_closed: %w", err)
			}
			p.IsClosed = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (p *Poll) GetID() (value int64) {
	if p == nil {
		return
	}
	return p.ID
}

// GetQuestion returns value of Question field.
func (p *Poll) GetQuestion() (value string) {
	if p == nil {
		return
	}
	return p.Question
}

// GetOptions returns value of Options field.
func (p *Poll) GetOptions() (value []PollOption) {
	if p == nil {
		return
	}
	return p.Options
}

// GetTotalVoterCount returns value of TotalVoterCount field.
func (p *Poll) GetTotalVoterCount() (value int32) {
	if p == nil {
		return
	}
	return p.TotalVoterCount
}

// GetRecentVoterIDs returns value of RecentVoterIDs field.
func (p *Poll) GetRecentVoterIDs() (value []MessageSenderClass) {
	if p == nil {
		return
	}
	return p.RecentVoterIDs
}

// GetIsAnonymous returns value of IsAnonymous field.
func (p *Poll) GetIsAnonymous() (value bool) {
	if p == nil {
		return
	}
	return p.IsAnonymous
}

// GetType returns value of Type field.
func (p *Poll) GetType() (value PollTypeClass) {
	if p == nil {
		return
	}
	return p.Type
}

// GetOpenPeriod returns value of OpenPeriod field.
func (p *Poll) GetOpenPeriod() (value int32) {
	if p == nil {
		return
	}
	return p.OpenPeriod
}

// GetCloseDate returns value of CloseDate field.
func (p *Poll) GetCloseDate() (value int32) {
	if p == nil {
		return
	}
	return p.CloseDate
}

// GetIsClosed returns value of IsClosed field.
func (p *Poll) GetIsClosed() (value bool) {
	if p == nil {
		return
	}
	return p.IsClosed
}
