# mackerel-plugin-cgroup-stats

## Description

Mackerel metrics plugin to get stats of cgroups v2.

## Synopsis

```sh
mackerel-plugin-cgroup-stats
```

## Installation

```sh
sudo mkr plugin install Arthur1/mackerel-plugin-cgroup-stats
```

## Setting for mackerel-agent

```
[plugin.metrics.cgroup-stats-hoge]
command = ["/opt/mackerel-agent/plugins/bin/mackerel-plugin-cgroup-stats", "-metrickey", "hoge", "-group", "hoge.service"]
```

## Usage

### Options

```
  -group string
        group name
  -metrickey string
        metric key (required)
  -slice string
        slice name (default "system.slice")
```
