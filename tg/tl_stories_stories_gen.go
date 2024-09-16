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

// StoriesStories represents TL type `stories.stories#63c3dd0a`.
// List of stories¹
//
// Links:
//  1. https://core.telegram.org/api/stories#pinned-or-archived-stories
//
// See https://core.telegram.org/constructor/stories.stories for reference.
type StoriesStories struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Total number of stories that can be fetched
	Count int
	// Stories
	Stories []StoryItemClass
	// IDs of pinned stories.
	//
	// Use SetPinnedToTop and GetPinnedToTop helpers.
	PinnedToTop []int
	// Mentioned chats
	Chats []ChatClass
	// Mentioned users
	Users []UserClass
}

// StoriesStoriesTypeID is TL type id of StoriesStories.
const StoriesStoriesTypeID = 0x63c3dd0a

// Ensuring interfaces in compile-time for StoriesStories.
var (
	_ bin.Encoder     = &StoriesStories{}
	_ bin.Decoder     = &StoriesStories{}
	_ bin.BareEncoder = &StoriesStories{}
	_ bin.BareDecoder = &StoriesStories{}
)

func (s *StoriesStories) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Count == 0) {
		return false
	}
	if !(s.Stories == nil) {
		return false
	}
	if !(s.PinnedToTop == nil) {
		return false
	}
	if !(s.Chats == nil) {
		return false
	}
	if !(s.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoriesStories) String() string {
	if s == nil {
		return "StoriesStories(nil)"
	}
	type Alias StoriesStories
	return fmt.Sprintf("StoriesStories%+v", Alias(*s))
}

// FillFrom fills StoriesStories from given interface.
func (s *StoriesStories) FillFrom(from interface {
	GetCount() (value int)
	GetStories() (value []StoryItemClass)
	GetPinnedToTop() (value []int, ok bool)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	s.Count = from.GetCount()
	s.Stories = from.GetStories()
	if val, ok := from.GetPinnedToTop(); ok {
		s.PinnedToTop = val
	}

	s.Chats = from.GetChats()
	s.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesStories) TypeID() uint32 {
	return StoriesStoriesTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesStories) TypeName() string {
	return "stories.stories"
}

// TypeInfo returns info about TL type.
func (s *StoriesStories) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.stories",
		ID:   StoriesStoriesTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Stories",
			SchemaName: "stories",
		},
		{
			Name:       "PinnedToTop",
			SchemaName: "pinned_to_top",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *StoriesStories) SetFlags() {
	if !(s.PinnedToTop == nil) {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *StoriesStories) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stories.stories#63c3dd0a as nil")
	}
	b.PutID(StoriesStoriesTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoriesStories) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stories.stories#63c3dd0a as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.stories#63c3dd0a: field flags: %w", err)
	}
	b.PutInt(s.Count)
	b.PutVectorHeader(len(s.Stories))
	for idx, v := range s.Stories {
		if v == nil {
			return fmt.Errorf("unable to encode stories.stories#63c3dd0a: field stories element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stories.stories#63c3dd0a: field stories element with index %d: %w", idx, err)
		}
	}
	if s.Flags.Has(0) {
		b.PutVectorHeader(len(s.PinnedToTop))
		for _, v := range s.PinnedToTop {
			b.PutInt(v)
		}
	}
	b.PutVectorHeader(len(s.Chats))
	for idx, v := range s.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode stories.stories#63c3dd0a: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stories.stories#63c3dd0a: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(s.Users))
	for idx, v := range s.Users {
		if v == nil {
			return fmt.Errorf("unable to encode stories.stories#63c3dd0a: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stories.stories#63c3dd0a: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoriesStories) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stories.stories#63c3dd0a to nil")
	}
	if err := b.ConsumeID(StoriesStoriesTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.stories#63c3dd0a: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoriesStories) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stories.stories#63c3dd0a to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stories.stories#63c3dd0a: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stories.stories#63c3dd0a: field count: %w", err)
		}
		s.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.stories#63c3dd0a: field stories: %w", err)
		}

		if headerLen > 0 {
			s.Stories = make([]StoryItemClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeStoryItem(b)
			if err != nil {
				return fmt.Errorf("unable to decode stories.stories#63c3dd0a: field stories: %w", err)
			}
			s.Stories = append(s.Stories, value)
		}
	}
	if s.Flags.Has(0) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.stories#63c3dd0a: field pinned_to_top: %w", err)
		}

		if headerLen > 0 {
			s.PinnedToTop = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode stories.stories#63c3dd0a: field pinned_to_top: %w", err)
			}
			s.PinnedToTop = append(s.PinnedToTop, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.stories#63c3dd0a: field chats: %w", err)
		}

		if headerLen > 0 {
			s.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode stories.stories#63c3dd0a: field chats: %w", err)
			}
			s.Chats = append(s.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.stories#63c3dd0a: field users: %w", err)
		}

		if headerLen > 0 {
			s.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode stories.stories#63c3dd0a: field users: %w", err)
			}
			s.Users = append(s.Users, value)
		}
	}
	return nil
}

// GetCount returns value of Count field.
func (s *StoriesStories) GetCount() (value int) {
	if s == nil {
		return
	}
	return s.Count
}

// GetStories returns value of Stories field.
func (s *StoriesStories) GetStories() (value []StoryItemClass) {
	if s == nil {
		return
	}
	return s.Stories
}

// SetPinnedToTop sets value of PinnedToTop conditional field.
func (s *StoriesStories) SetPinnedToTop(value []int) {
	s.Flags.Set(0)
	s.PinnedToTop = value
}

// GetPinnedToTop returns value of PinnedToTop conditional field and
// boolean which is true if field was set.
func (s *StoriesStories) GetPinnedToTop() (value []int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.PinnedToTop, true
}

// GetChats returns value of Chats field.
func (s *StoriesStories) GetChats() (value []ChatClass) {
	if s == nil {
		return
	}
	return s.Chats
}

// GetUsers returns value of Users field.
func (s *StoriesStories) GetUsers() (value []UserClass) {
	if s == nil {
		return
	}
	return s.Users
}

// MapStories returns field Stories wrapped in StoryItemClassArray helper.
func (s *StoriesStories) MapStories() (value StoryItemClassArray) {
	return StoryItemClassArray(s.Stories)
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (s *StoriesStories) MapChats() (value ChatClassArray) {
	return ChatClassArray(s.Chats)
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (s *StoriesStories) MapUsers() (value UserClassArray) {
	return UserClassArray(s.Users)
}
