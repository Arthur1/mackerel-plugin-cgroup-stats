package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/containerd/cgroups/v3/cgroup2"
	mp "github.com/mackerelio/go-mackerel-plugin"
)

func main() {
	plugin := new(Plugin)

	flag.StringVar(&plugin.metricKey, "metrickey", "", "metric key (required)")
	flag.StringVar(&plugin.slice, "slice", "system.slice", "slice name")
	flag.StringVar(&plugin.group, "group", "", "group name")
	flag.Parse()

	if plugin.metricKey == "" {
		log.Fatalln("-metrickey is required")
	}

	mackerelPlugin := mp.NewMackerelPlugin(plugin)
	mackerelPlugin.Run()
}

type Plugin struct {
	metricKey string
	slice     string
	group     string
}

var _ mp.PluginWithPrefix = new(Plugin)

func (p *Plugin) FetchMetrics() (map[string]float64, error) {
	m, err := cgroup2.LoadSystemd(p.slice, p.group)
	if err != nil {
		return nil, err
	}
	stats, err := m.Stat()
	if err != nil {
		return nil, err
	}
	metrics := make(map[string]float64, 2)
	metrics[fmt.Sprintf("memory.%s.usage", p.metricKey)] = float64(stats.Memory.Usage)
	metrics[fmt.Sprintf("memory.%s.limit", p.metricKey)] = float64(stats.Memory.UsageLimit)
	// TODO: implements other metrics
	return metrics, nil
}

func (p *Plugin) GraphDefinition() map[string]mp.Graphs {
	return map[string]mp.Graphs{
		"memory.#": {
			Label: "Cgroups Memory",
			Unit:  mp.UnitBytes,
			Metrics: []mp.Metrics{
				{Name: "limit", Label: "Limit"},
				{Name: "usage", Label: "Usage"},
			},
		},
	}
}

func (p *Plugin) MetricKeyPrefix() string {
	return "cgroup-stats"
}
