package message //nolint:dupl

import (
	"fmt"

	"github.com/bhaney/rtsp-simple-server/internal/rtmp/chunk"
	"github.com/bhaney/rtsp-simple-server/internal/rtmp/rawmessage"
)

// MsgUserControlPingRequest is a user control message.
type MsgUserControlPingRequest struct {
	ServerTime uint32
}

// Unmarshal implements Message.
func (m *MsgUserControlPingRequest) Unmarshal(raw *rawmessage.Message) error {
	if raw.ChunkStreamID != ControlChunkStreamID {
		return fmt.Errorf("unexpected chunk stream ID")
	}

	if len(raw.Body) != 6 {
		return fmt.Errorf("invalid body size")
	}

	m.ServerTime = uint32(raw.Body[2])<<24 | uint32(raw.Body[3])<<16 | uint32(raw.Body[4])<<8 | uint32(raw.Body[5])

	return nil
}

// Marshal implements Message.
func (m MsgUserControlPingRequest) Marshal() (*rawmessage.Message, error) {
	buf := make([]byte, 6)

	buf[0] = byte(UserControlTypePingRequest >> 8)
	buf[1] = byte(UserControlTypePingRequest)
	buf[2] = byte(m.ServerTime >> 24)
	buf[3] = byte(m.ServerTime >> 16)
	buf[4] = byte(m.ServerTime >> 8)
	buf[5] = byte(m.ServerTime)

	return &rawmessage.Message{
		ChunkStreamID: ControlChunkStreamID,
		Type:          chunk.MessageTypeUserControl,
		Body:          buf,
	}, nil
}
