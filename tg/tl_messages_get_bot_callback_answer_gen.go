// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// MessagesGetBotCallbackAnswerRequest represents TL type `messages.getBotCallbackAnswer#9342ca07`.
// Press an inline callback button and get a callback answer from the bot
//
// See https://core.telegram.org/method/messages.getBotCallbackAnswer for reference.
type MessagesGetBotCallbackAnswerRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this is a "play game" button
	Game bool
	// Where was the inline keyboard sent
	Peer InputPeerClass
	// ID of the Message with the inline keyboard
	MsgID int
	// Callback data
	//
	// Use SetData and GetData helpers.
	Data []byte
	// For buttons requiring you to verify your identity with your 2FA password¹, the SRP payload generated using SRP².
	//
	// Links:
	//  1) https://core.telegram.org/constructor/keyboardButtonCallback
	//  2) https://core.telegram.org/api/srp
	//
	// Use SetPassword and GetPassword helpers.
	Password InputCheckPasswordSRPClass
}

// MessagesGetBotCallbackAnswerRequestTypeID is TL type id of MessagesGetBotCallbackAnswerRequest.
const MessagesGetBotCallbackAnswerRequestTypeID = 0x9342ca07

func (g *MessagesGetBotCallbackAnswerRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Game == false) {
		return false
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.MsgID == 0) {
		return false
	}
	if !(g.Data == nil) {
		return false
	}
	if !(g.Password == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetBotCallbackAnswerRequest) String() string {
	if g == nil {
		return "MessagesGetBotCallbackAnswerRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesGetBotCallbackAnswerRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(g.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(g.Peer))
	sb.WriteString(",\n")
	sb.WriteString("\tMsgID: ")
	sb.WriteString(fmt.Sprint(g.MsgID))
	sb.WriteString(",\n")
	if g.Flags.Has(0) {
		sb.WriteString("\tData: ")
		sb.WriteString(fmt.Sprint(g.Data))
		sb.WriteString(",\n")
	}
	if g.Flags.Has(2) {
		sb.WriteString("\tPassword: ")
		sb.WriteString(fmt.Sprint(g.Password))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *MessagesGetBotCallbackAnswerRequest) TypeID() uint32 {
	return MessagesGetBotCallbackAnswerRequestTypeID
}

// Encode implements bin.Encoder.
func (g *MessagesGetBotCallbackAnswerRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getBotCallbackAnswer#9342ca07 as nil")
	}
	b.PutID(MessagesGetBotCallbackAnswerRequestTypeID)
	if !(g.Game == false) {
		g.Flags.Set(1)
	}
	if !(g.Data == nil) {
		g.Flags.Set(0)
	}
	if !(g.Password == nil) {
		g.Flags.Set(2)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getBotCallbackAnswer#9342ca07: field flags: %w", err)
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getBotCallbackAnswer#9342ca07: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getBotCallbackAnswer#9342ca07: field peer: %w", err)
	}
	b.PutInt(g.MsgID)
	if g.Flags.Has(0) {
		b.PutBytes(g.Data)
	}
	if g.Flags.Has(2) {
		if g.Password == nil {
			return fmt.Errorf("unable to encode messages.getBotCallbackAnswer#9342ca07: field password is nil")
		}
		if err := g.Password.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.getBotCallbackAnswer#9342ca07: field password: %w", err)
		}
	}
	return nil
}

// SetGame sets value of Game conditional field.
func (g *MessagesGetBotCallbackAnswerRequest) SetGame(value bool) {
	if value {
		g.Flags.Set(1)
		g.Game = true
	} else {
		g.Flags.Unset(1)
		g.Game = false
	}
}

// GetGame returns value of Game conditional field.
func (g *MessagesGetBotCallbackAnswerRequest) GetGame() (value bool) {
	return g.Flags.Has(1)
}

// GetPeer returns value of Peer field.
func (g *MessagesGetBotCallbackAnswerRequest) GetPeer() (value InputPeerClass) {
	return g.Peer
}

// GetMsgID returns value of MsgID field.
func (g *MessagesGetBotCallbackAnswerRequest) GetMsgID() (value int) {
	return g.MsgID
}

// SetData sets value of Data conditional field.
func (g *MessagesGetBotCallbackAnswerRequest) SetData(value []byte) {
	g.Flags.Set(0)
	g.Data = value
}

// GetData returns value of Data conditional field and
// boolean which is true if field was set.
func (g *MessagesGetBotCallbackAnswerRequest) GetData() (value []byte, ok bool) {
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.Data, true
}

// SetPassword sets value of Password conditional field.
func (g *MessagesGetBotCallbackAnswerRequest) SetPassword(value InputCheckPasswordSRPClass) {
	g.Flags.Set(2)
	g.Password = value
}

// GetPassword returns value of Password conditional field and
// boolean which is true if field was set.
func (g *MessagesGetBotCallbackAnswerRequest) GetPassword() (value InputCheckPasswordSRPClass, ok bool) {
	if !g.Flags.Has(2) {
		return value, false
	}
	return g.Password, true
}

// Decode implements bin.Decoder.
func (g *MessagesGetBotCallbackAnswerRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getBotCallbackAnswer#9342ca07 to nil")
	}
	if err := b.ConsumeID(MessagesGetBotCallbackAnswerRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getBotCallbackAnswer#9342ca07: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.getBotCallbackAnswer#9342ca07: field flags: %w", err)
		}
	}
	g.Game = g.Flags.Has(1)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getBotCallbackAnswer#9342ca07: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getBotCallbackAnswer#9342ca07: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	if g.Flags.Has(0) {
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getBotCallbackAnswer#9342ca07: field data: %w", err)
		}
		g.Data = value
	}
	if g.Flags.Has(2) {
		value, err := DecodeInputCheckPasswordSRP(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getBotCallbackAnswer#9342ca07: field password: %w", err)
		}
		g.Password = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetBotCallbackAnswerRequest.
var (
	_ bin.Encoder = &MessagesGetBotCallbackAnswerRequest{}
	_ bin.Decoder = &MessagesGetBotCallbackAnswerRequest{}
)

// MessagesGetBotCallbackAnswer invokes method messages.getBotCallbackAnswer#9342ca07 returning error if any.
// Press an inline callback button and get a callback answer from the bot
//
// Possible errors:
//  400 BOT_RESPONSE_TIMEOUT: A timeout occurred while fetching data from the bot
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 DATA_INVALID: Encrypted data invalid
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  -503 Timeout: Timeout while fetching data
//
// See https://core.telegram.org/method/messages.getBotCallbackAnswer for reference.
func (c *Client) MessagesGetBotCallbackAnswer(ctx context.Context, request *MessagesGetBotCallbackAnswerRequest) (*MessagesBotCallbackAnswer, error) {
	var result MessagesBotCallbackAnswer

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
