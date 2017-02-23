package api

import "bytes"

// ExplicitTxCommand is an Explicit Addressing Command XBee frame.
type ExplicitTxCommand struct {
	// FrameID identifies the data frame for the host to correlate with a
	// subsequent ACK. If set to 0, the device does not send a response.
	FrameID uint8
	// Addr64 is tthe 64-bit address of the destination device. Reserved 64-bit
	// addresses:
	//
	// coordinator = 0x0000000000000000
	// Broadcast = 0x000000000000FFFF.
	Addr64 uint64
	// Addr16 is the 16-bit address of the destination device, if known. Set to
	// 0xFFFE if the address is unknown, or if sending a broadcast.
	Addr16 uint16
	// SrcEp is the source Endpoint for the transmission.
	SrcEp uint8
	// DstEp is the destination Endpoint for the transmission.
	DstEp uint8
	// ClusterID used in the transmission.
	ClusterID uint16
	// ProfileID used in the transmission.
	ProfileID uint16
	// BroadcastRadius is the maximum number of hops a broadcast transmission can
	// traverse. If set to 0, the device sets the transmission radius to the
	// network maximum hops value.
	BroadcastRadius uint8
	// Options is a bitfield of supported transmission options.
	Options byte
	// Data is the data payload sent to the destination device, up to NP bytes.
	Data []byte
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
	if _, err := buf.Write(uint64ToBytes(c.Addr64)); err != nil {
		return nil, err
	}
	if _, err := buf.Write(uint16ToBytes(c.Addr16)); err != nil {
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
