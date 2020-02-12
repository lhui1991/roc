package rocserv

import (
	"gitlab.pri.ibanyu.com/middleware/seaweed/xstat/xmetric"
	xprom "gitlab.pri.ibanyu.com/middleware/seaweed/xstat/xmetric/xprometheus"
)

const (
	namespacePalfish     = "palfish"
	namespaceAPM         = "apm"
	serverDurationSecond = "server_duration_second"
	serverRequestTotal   = "server_request_total"
	// client means that the duration is count from client side,
	//   which includes the network round-trip time and the server
	//   processing time.
	clientRequestTotal    = "client_request_total"
	clientRequestDuration = "client_request_duration"

	labelStatus = "status"

	apiType = "api"
)

var (
	buckets = []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}
	// 目前只用作sla统计，后续通过修改标签作为所有微服务的耗时统计
	_metricRequestDuration = xprom.NewHistogram(&xprom.HistogramVecOpts{
		Namespace:  namespacePalfish,
		Name:       serverDurationSecond,
		Help:       "sla calc request duration in seconds.",
		Buckets:    buckets,
		LabelNames: []string{xprom.LabelServiceName, xprom.LabelServiceID, xprom.LabelInstance, xprom.LabelAPI, xprom.LabelSource, xprom.LabelType},
	})
	_metricRequestTotal = xprom.NewCounter(&xprom.CounterVecOpts{
		Namespace:  namespacePalfish,
		Name:       serverRequestTotal,
		Help:       "sla calc request total.",
		LabelNames: []string{xprom.LabelServiceName, xprom.LabelServiceID, xprom.LabelInstance, xprom.LabelAPI, xprom.LabelSource, xprom.LabelType, labelStatus},
	})

	// apm 打点
	_metricAPMRequestDuration = xprom.NewHistogram(&xprom.HistogramVecOpts{
		Namespace:  namespaceAPM,
		Name:       clientRequestDuration,
		Help:       "apm client side request duration in seconds",
		Buckets:    buckets,
		LabelNames: []string{xprom.LabelCallerService, xprom.LabelCalleeService, xprom.LabelCallerEndpoint, xprom.LabelCalleeEndpoint, xprom.LabelCallerServiceID},
	})

	_metricAPMRequestTotal = xprom.NewCounter(&xprom.CounterVecOpts{
		Namespace:  namespaceAPM,
		Name:       clientRequestTotal,
		Help:       "apm client side request total",
		LabelNames: []string{xprom.LabelCallerService, xprom.LabelCalleeService, xprom.LabelCallerEndpoint, xprom.LabelCalleeEndpoint, xprom.LabelCallerServiceID, xprom.LabelCallStatus},
	})

	_metricAPIRequestCount = xprom.NewCounter(&xprom.CounterVecOpts{
		Namespace:  namespacePalfish,
		Subsystem:  apiType,
		Name:       "request_count",
		Help:       "api request count",
		LabelNames: []string{xprom.LabelGroupName, xprom.LabelServiceName, xprom.LabelAPI},
	})

	_metricAPIRequestTime = xprom.NewHistogram(&xprom.HistogramVecOpts{
		Namespace:  namespacePalfish,
		Subsystem:  apiType,
		Name:       "request_duration",
		Buckets:    []float64{10, 50, 100, 200, 300, 500, 1000, 3000, 5000, 10000},
		Help:       "api request duration in millisecond",
		LabelNames: []string{xprom.LabelGroupName, xprom.LabelServiceName, xprom.LabelAPI},
	})
)

func GetSlaDurationMetric() xmetric.Histogram {
	return _metricRequestDuration
}
func GetSlaRequestTotalMetric() xmetric.Counter {
	return _metricRequestTotal
}

// GetAPIRequestCountMetric 解决server/go/util/servbase/monitor与roc循环引用及重名metric的问题
func GetAPIRequestCountMetric() xmetric.Counter {
	return _metricAPIRequestCount
}

// GetAPIRequestTimeMetric 解决server/go/util/servbase/monitor与roc循环引用及重名metric的问题
func GetAPIRequestTimeMetric() xmetric.Histogram {
	return _metricAPIRequestTime
}
