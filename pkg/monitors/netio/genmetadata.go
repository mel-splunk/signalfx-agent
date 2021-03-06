// Code generated by monitor-code-gen. DO NOT EDIT.

package netio

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "net-io"

var groupSet = map[string]bool{}

const (
	ifErrorsRx   = "if_errors.rx"
	ifErrorsTx   = "if_errors.tx"
	ifOctetsRx   = "if_octets.rx"
	ifOctetsTx   = "if_octets.tx"
	ifPacketsRx  = "if_packets.rx"
	ifPacketsTx  = "if_packets.tx"
	networkTotal = "network.total"
)

var metricSet = map[string]monitors.MetricInfo{
	ifErrorsRx:   {Type: datapoint.Counter},
	ifErrorsTx:   {Type: datapoint.Counter},
	ifOctetsRx:   {Type: datapoint.Counter},
	ifOctetsTx:   {Type: datapoint.Counter},
	ifPacketsRx:  {Type: datapoint.Counter},
	ifPacketsTx:  {Type: datapoint.Counter},
	networkTotal: {Type: datapoint.Counter},
}

var defaultMetrics = map[string]bool{
	ifErrorsRx:   true,
	ifErrorsTx:   true,
	ifOctetsRx:   true,
	ifOctetsTx:   true,
	networkTotal: true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:       "net-io",
	DefaultMetrics:    defaultMetrics,
	Metrics:           metricSet,
	MetricsExhaustive: false,
	Groups:            groupSet,
	GroupMetricsMap:   groupMetricsMap,
	SendAll:           false,
}
