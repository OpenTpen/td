//go:build !no_gotd_slices
// +build !no_gotd_slices

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

// PaidReactionPrivacyClassArray is adapter for slice of PaidReactionPrivacyClass.
type PaidReactionPrivacyClassArray []PaidReactionPrivacyClass

// Sort sorts slice of PaidReactionPrivacyClass.
func (s PaidReactionPrivacyClassArray) Sort(less func(a, b PaidReactionPrivacyClass) bool) PaidReactionPrivacyClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PaidReactionPrivacyClass.
func (s PaidReactionPrivacyClassArray) SortStable(less func(a, b PaidReactionPrivacyClass) bool) PaidReactionPrivacyClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PaidReactionPrivacyClass.
func (s PaidReactionPrivacyClassArray) Retain(keep func(x PaidReactionPrivacyClass) bool) PaidReactionPrivacyClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s PaidReactionPrivacyClassArray) First() (v PaidReactionPrivacyClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PaidReactionPrivacyClassArray) Last() (v PaidReactionPrivacyClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PaidReactionPrivacyClassArray) PopFirst() (v PaidReactionPrivacyClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PaidReactionPrivacyClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PaidReactionPrivacyClassArray) Pop() (v PaidReactionPrivacyClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsPaidReactionPrivacyPeer returns copy with only PaidReactionPrivacyPeer constructors.
func (s PaidReactionPrivacyClassArray) AsPaidReactionPrivacyPeer() (to PaidReactionPrivacyPeerArray) {
	for _, elem := range s {
		value, ok := elem.(*PaidReactionPrivacyPeer)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// PaidReactionPrivacyPeerArray is adapter for slice of PaidReactionPrivacyPeer.
type PaidReactionPrivacyPeerArray []PaidReactionPrivacyPeer

// Sort sorts slice of PaidReactionPrivacyPeer.
func (s PaidReactionPrivacyPeerArray) Sort(less func(a, b PaidReactionPrivacyPeer) bool) PaidReactionPrivacyPeerArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PaidReactionPrivacyPeer.
func (s PaidReactionPrivacyPeerArray) SortStable(less func(a, b PaidReactionPrivacyPeer) bool) PaidReactionPrivacyPeerArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PaidReactionPrivacyPeer.
func (s PaidReactionPrivacyPeerArray) Retain(keep func(x PaidReactionPrivacyPeer) bool) PaidReactionPrivacyPeerArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s PaidReactionPrivacyPeerArray) First() (v PaidReactionPrivacyPeer, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PaidReactionPrivacyPeerArray) Last() (v PaidReactionPrivacyPeer, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PaidReactionPrivacyPeerArray) PopFirst() (v PaidReactionPrivacyPeer, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PaidReactionPrivacyPeer
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PaidReactionPrivacyPeerArray) Pop() (v PaidReactionPrivacyPeer, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
