package api

// ExplicitRxIndicator is received when a device configured with explicit API Rx
// Indicator (AO = 1) receives an RF packet.
type ExplicitRxIndicator struct {
	DstAddr64 uint64
	DstAddr16 uint16
	SrcEp     uint8
	DstEp     uint8
	ClusterID uint16
	ProfileID uint16
	Options   byte
	Data      []byte
}
