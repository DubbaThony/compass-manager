package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

type Metrics struct {
	configured   prometheus.Counter
	registered   prometheus.Counter
	unregistered prometheus.Counter
}

func NewMetrics() Metrics {
	m := Metrics{
		configured: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "configure_counter",
			Help: "Number of cluster configured",
		}),
		registered: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "register_counter",
			Help: "Number of cluster registered",
		}),
		unregistered: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "unregister_counter",
			Help: "Number of cluster unregistered",
		}),
	}
	metrics.Registry.MustRegister(m.configured, m.registered, m.unregistered)
	return m
}

func (m Metrics) IncConfigure() {
	m.configured.Inc()
}

func (m Metrics) IncRegister() {
	m.registered.Inc()
}

func (m Metrics) IncUnregister() {
	m.unregistered.Inc()
}