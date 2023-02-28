package metricsfilter

import (
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/prometheus/common/model"
)

type MetricsAllowed struct {
	items hashset.Set
	count int
}

func NewMetricAllowed() *MetricsAllowed {
	m := MetricsAllowed{items: *hashset.New(), count: 0}
	return &m
}

func (m *MetricsAllowed) Update(discovered *model.LabelValues, included *[]string, excluded *[]string) *MetricsAllowed {
	// cleanup
	m.items.Clear()

	// add discoverd items
	for _, item := range *discovered {
		m.items.Add(string(item))
	}

	// add statically included items
	for _, item := range *included {
		m.items.Add(string(item))
	}

	// delete statically excluded items
	for _, item := range *excluded {
		m.items.Remove(string(item))
	}
	return m
}

func (m *MetricsAllowed) Contains(metricName string) bool {
	return m.items.Contains(metricName)
}
