package api

// ExplicitRxIndicator is received when a device configured with explicit API Rx
// Indicator (AO = 1) receives an RF packet.
type ExplicitRxIndicator struct {
	// Addr64 is the sender's 64-bit address. Set to 0xFFFFFFFFFFFFFFFF (unknown
	// 64-bit address) if the sender's 64-bit address is unknown.
	Addr64 uint64
	// Addr16 is the sender's 16-bit address.
	Addr16 uint16
	// SrcEp is the endpoint of the source that initiates transmission.
	SrcEp uint8
	// DstEp is the endpoint of the destination that the message is addressed to.
	DstEp uint8
	// ClusterID that the frame is addressed to.
	ClusterID uint16
	// ProfileID that the fame is addressed to.
	ProfileID uint16
	// Options are the receive options.
	Options byte
	// Data is the received RF data.
	Data []byte
}
