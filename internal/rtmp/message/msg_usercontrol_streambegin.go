package message //nolint:dupl

import (
	"fmt"

	"github.com/bhaney/rtsp-simple-server/internal/rtmp/chunk"
	"github.com/bhaney/rtsp-simple-server/internal/rtmp/rawmessage"
)

// MsgUserControlStreamBegin is a user control message.
type MsgUserControlStreamBegin struct {
	StreamID uint32
}

// Unmarshal implements Message.
func (m *MsgUserControlStreamBegin) Unmarshal(raw *rawmessage.Message) error {
	if raw.ChunkStreamID != ControlChunkStreamID {
		return fmt.Errorf("unexpected chunk stream ID")
	}

	if len(raw.Body) != 6 {
		return fmt.Errorf("invalid body size")
	}

	m.StreamID = uint32(raw.Body[2])<<24 | uint32(raw.Body[3])<<16 | uint32(raw.Body[4])<<8 | uint32(raw.Body[5])

	return nil
}

// Marshal implements Message.
func (m MsgUserControlStreamBegin) Marshal() (*rawmessage.Message, error) {
	buf := make([]byte, 6)

	buf[0] = byte(UserControlTypeStreamBegin >> 8)
	buf[1] = byte(UserControlTypeStreamBegin)
	buf[2] = byte(m.StreamID >> 24)
	buf[3] = byte(m.StreamID >> 16)
	buf[4] = byte(m.StreamID >> 8)
	buf[5] = byte(m.StreamID)

	return &rawmessage.Message{
		ChunkStreamID: ControlChunkStreamID,
		Type:          chunk.MessageTypeUserControl,
		Body:          buf,
	}, nil
}
