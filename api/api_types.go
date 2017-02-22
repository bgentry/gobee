package api

type apiTypeID byte

const (
	apiTypeAT         apiTypeID = 0x08
	apiTypeExplicitTx apiTypeID = 0x11
	apiTypeExplicitRx apiTypeID = 0x91
)

// TxDeliveryStatus describes the possible delivery status results.
type TxDeliveryStatus byte

const (
	// TxDeliverySuccess is a successful delivery.
	TxDeliverySuccess TxDeliveryStatus = 0x00
	// TxDeliveryMACAckFailure ...
	TxDeliveryMACAckFailure TxDeliveryStatus = 0x01
	// TxDeliveryCCAFailure ...
	TxDeliveryCCAFailure TxDeliveryStatus = 0x02
	// TxDeliveryInvalidDestinationEndpoint ...
	TxDeliveryInvalidDestinationEndpoint TxDeliveryStatus = 0x15
	// TxDeliveryNetworkAckFailure ...
	TxDeliveryNetworkAckFailure TxDeliveryStatus = 0x21
	// TxDeliveryNotJoinedToNetwork ...
	TxDeliveryNotJoinedToNetwork TxDeliveryStatus = 0x22
	// TxDeliverySelfAddressed indicates that the frame was addressed to its
	// sender and was delivered.
	TxDeliverySelfAddressed TxDeliveryStatus = 0x23
	// TxDeliveryAddressNotFound ...
	TxDeliveryAddressNotFound TxDeliveryStatus = 0x24
	// TxDeliveryRouteNotFound ...
	TxDeliveryRouteNotFound TxDeliveryStatus = 0x25
	// TxDeliveryBroadcastSourceFailedToHearRelay ...
	TxDeliveryBroadcastSourceFailedToHearRelay TxDeliveryStatus = 0x26
	// TxDeliveryInvalidBindingTableIndex is returned when the binding table index is invalid.
	TxDeliveryInvalidBindingTableIndex TxDeliveryStatus = 0x2B
	// TxDeliveryResourceError signifies a lack of free buffers, timers, etc.
	TxDeliveryResourceError TxDeliveryStatus = 0x2C
	// TxDeliveryAttemptedBroadcastWithAPSTx is returned when a broadcast is
	// attempted with APS transmission
	TxDeliveryAttemptedBroadcastWithAPSTx TxDeliveryStatus = 0x2D
	// TxDeliveryAttemptedUnicast is from an attempted unicast with APS
	// transmission, but EE=0.
	TxDeliveryAttemptedUnicast TxDeliveryStatus = 0x2E
	// TxDeliveryResourceError32 signifies a lack of free buffers, timers, etc.
	TxDeliveryResourceError32 TxDeliveryStatus = 0x32
	// TxDeliveryPayloadTooLarge means the data payload was too large.
	TxDeliveryPayloadTooLarge TxDeliveryStatus = 0x74
	// TxDeliveryIndirectMessageUnrequested = Indirect message unrequested
	TxDeliveryIndirectMessageUnrequested TxDeliveryStatus = 0x75
)

// TxDiscoveryStatus ...
type TxDiscoveryStatus byte

const (
	// TxDiscoveryNoDiscoveryOverhead ...
	TxDiscoveryNoDiscoveryOverhead TxDiscoveryStatus = 0x01
	// TxDiscoveryRouteDiscovery ...
	TxDiscoveryRouteDiscovery TxDiscoveryStatus = 0x02
	// TxDiscoveryAddressAndRoute ...
	TxDiscoveryAddressAndRoute TxDiscoveryStatus = 0x03
	// TxDiscoveryExtendedTimeoutDiscovery ...
	TxDiscoveryExtendedTimeoutDiscovery TxDiscoveryStatus = 0x40
)
