package api

import "bytes"

// ATCommand is an AT Command XBee frame.
type ATCommand struct {
	FrameID   uint8
	Type      [2]byte
	Parameter []byte
}

// Bytes turns the ATCommand frame into bytes
func (c *ATCommand) Bytes() ([]byte, error) {
	var buf bytes.Buffer
	if err := buf.WriteByte(byte(apiTypeAT)); err != nil {
		return nil, err
	}
	if err := buf.WriteByte(c.FrameID); err != nil {
		return nil, err
	}
	if _, err := buf.Write(c.Type[:]); err != nil {
		return nil, err
	}
	if _, err := buf.Write(c.Parameter); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ATCommandStatus represents the possible status results of an AT Command.
type ATCommandStatus byte

const (
	ATCommandStatusOK               ATCommandStatus = 0
	ATCommandStatusError            ATCommandStatus = 1
	ATCommandStatusInvalidCommand   ATCommandStatus = 2
	ATCommandStatusInvalidParameter ATCommandStatus = 3
	ATCommandStatusTxFailure        ATCommandStatus = 4
)
