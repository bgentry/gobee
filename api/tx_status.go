package api

// TxStatus is received when a TxCommand or ExplicitTxCommand completes.
type TxStatus struct {
	FrameID   uint8
	Addr16    uint16
	Retries   uint8
	Delivery  TxDeliveryStatus
	Discovery TxDiscoveryStatus
}
