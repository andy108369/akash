package hostname_operator

import (
	ctypes "github.com/ovrclk/akash/provider/cluster/types"
	mtypes "github.com/ovrclk/akash/x/market/types/v1beta2"
	"time"
)

type managedHostname struct {
	lastEvent    ctypes.HostnameResourceEvent
	presentLease mtypes.LeaseID

	presentServiceName  string
	presentExternalPort uint32
	lastChangeAt        time.Time
}

type hostnameOperatorConfig struct {
	listenAddress        string
	pruneInterval        time.Duration
	ignoreListEntryLimit uint
	ignoreListAgeLimit   time.Duration
	webRefreshInterval   time.Duration
	retryDelay           time.Duration
	eventFailureLimit    uint
}