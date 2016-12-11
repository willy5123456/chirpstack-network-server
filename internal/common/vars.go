package common

import (
	"time"

	"github.com/brocaar/lorawan/band"
)

// TODO: move these vars to Context?

// NodeTXPayloadQueueTTL defines the TTL of the node TXPayload queue
var NodeTXPayloadQueueTTL = time.Hour * 24 * 5

// NodeSessionTTL defines the TTL of a node session (will be renewed on each
// activity)
var NodeSessionTTL = time.Hour * 24 * 5

// MACPendingTTL defines the TTL of a pending mac-command.
var MACPendingTTL = time.Hour * 24 * 5

// Band is the ISM band configuration to use
var Band band.Band

// BandName is the name of the used ISM band
var BandName band.Name
