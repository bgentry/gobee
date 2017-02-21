package api

import "bytes"

// ExplicitTxCommand is an Explicit Addressing Command XBee frame.
type ExplicitTxCommand struct {
	FrameID         uint8
	DstAddr64       uint64
	DstAddr16       uint16
	SrcEp           uint8
	DstEp           uint8
	ClusterID       uint16
	ProfileID       uint16
	BroadcastRadius uint8
	Options         byte
	Data            []byte
}

// Bytes turns the ExplicitTxCommand frame into bytes
func (c *ExplicitTxCommand) Bytes() ([]byte, error) {
	var buf bytes.Buffer
	if err := buf.WriteByte(byte(apiTypeExplicitTx)); err != nil {
		return nil, err
	}
	if err := buf.WriteByte(c.FrameID); err != nil {
		return nil, err
	}
	if _, err := buf.Write(uint64ToBytes(c.DstAddr64)); err != nil {
		return nil, err
	}
	if _, err := buf.Write(uint16ToBytes(c.DstAddr16)); err != nil {
		return nil, err
	}
	if err := buf.WriteByte(c.SrcEp); err != nil {
		return nil, err
	}
	if err := buf.WriteByte(c.DstEp); err != nil {
		return nil, err
	}
	if _, err := buf.Write(uint16ToBytes(c.ClusterID)); err != nil {
		return nil, err
	}
	if _, err := buf.Write(uint16ToBytes(c.ProfileID)); err != nil {
		return nil, err
	}
	if err := buf.WriteByte(c.BroadcastRadius); err != nil {
		return nil, err
	}
	if err := buf.WriteByte(c.Options); err != nil {
		return nil, err
	}
	if _, err := buf.Write(c.Data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
