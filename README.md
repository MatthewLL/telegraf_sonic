# Telegraf [![Circle CI](https://circleci.com/gh/influxdata/telegraf.svg?style=svg)](https://circleci.com/gh/influxdata/telegraf) [![Docker pulls](https://img.shields.io/docker/pulls/library/telegraf.svg)](https://hub.docker.com/_/telegraf/)
[![Slack Status](https://img.shields.io/badge/slack-join_chat-white.svg?logo=slack&style=social)](https://www.influxdata.com/slack)

Telegraf is an agent for collecting, processing, aggregating, and writing metrics.

Design goals are to have a minimal memory footprint with a plugin system so
that developers in the community can easily add support for collecting
metrics.

Telegraf is plugin-driven and has the concept of 4 distinct plugin types:

1. [Input Plugins](#input-plugins) collect metrics from the system, services, or 3rd party APIs
2. [Processor Plugins](#processor-plugins) transform, decorate, and/or filter metrics
3. [Aggregator Plugins](#aggregator-plugins) create aggregate metrics (e.g. mean, min, max, quantiles, etc.)
4. [Output Plugins](#output-plugins) write metrics to various destinations

New plugins are designed to be easy to contribute, pull requests are welcomed
and we work to incorporate as many pull requests as possible.

## Try in Browser :rocket:

You can try Telegraf right in your browser in the [Telegraf playground](https://rootnroll.com/d/telegraf/).

## Contributing

There are many ways to contribute:
- Fix and [report bugs](https://github.com/influxdata/telegraf/issues/new)
- [Improve documentation](https://github.com/influxdata/telegraf/issues?q=is%3Aopen+label%3Adocumentation)
- [Review code and feature proposals](https://github.com/influxdata/telegraf/pulls)
- Answer questions and discuss here on github and on the [Community Site](https://community.influxdata.com/)
- [Contribute plugins](CONTRIBUTING.md)
- [Contribute external plugins](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/execd/shim) *(experimental)*

## Minimum Requirements

Telegraf shares the same [minimum requirements][] as Go:
- Linux kernel version 2.6.23 or later
- Windows 7 or later
- FreeBSD 11.2 or later
- MacOS 10.11 El Capitan or later

[minimum requirements]: https://github.com/golang/go/wiki/MinimumRequirements#minimum-requirements

## Installation:

You can download the binaries directly from the [downloads](https://www.influxdata.com/downloads) page
or from the [releases](https://github.com/influxdata/telegraf/releases) section.

### Ansible Role:

Ansible role: https://github.com/rossmcdonald/telegraf

### From Source:

Telegraf requires Go version 1.13 or newer, the Makefile requires GNU make.

1. [Install Go](https://golang.org/doc/install) >=1.13 (1.15 recommended)
2. Clone the Telegraf repository:
   ```
   cd ~/src
   git clone https://github.com/influxdata/telegraf.git
   ```
3. Run `make` from the source directory
   ```
   cd ~/src/telegraf
   make
   ```

### Changelog

View the [changelog](/CHANGELOG.md) for the latest updates and changes by
version.

### Nightly Builds

These builds are generated from the master branch:
- [telegraf-nightly_darwin_amd64.tar.gz](https://dl.influxdata.com/telegraf/nightlies/telegraf-nightly_darwin_amd64.tar.gz)
- [telegraf_nightly_amd64.deb](https://dl.influxdata.com/telegraf/nightlies/telegraf_nightly_amd64.deb)
- [telegraf_nightly_arm64.deb](https://dl.influxdata.com/telegraf/nightlies/telegraf_nightly_arm64.deb)
- [telegraf-nightly.arm64.rpm](https://dl.influxdata.com/telegraf/nightlies/telegraf-nightly.arm64.rpm)
- [telegraf_nightly_armel.deb](https://dl.influxdata.com/telegraf/nightlies/telegraf_nightly_armel.deb)
- [telegraf-nightly.armel.rpm](https://dl.influxdata.com/telegraf/nightlies/telegraf-nightly.armel.rpm)
- [telegraf_nightly_armhf.deb](https://dl.influxdata.com/telegraf/nightlies/telegraf_nightly_armhf.deb)
- [telegraf-nightly.armv6hl.rpm](https://dl.influxdata.com/telegraf/nightlies/telegraf-nightly.armv6hl.rpm)
- [telegraf-nightly_freebsd_amd64.tar.gz](https://dl.influxdata.com/telegraf/nightlies/telegraf-nightly_freebsd_amd64.tar.gz)
- [telegraf-nightly_freebsd_i386.tar.gz](https://dl.influxdata.com/telegraf/nightlies/telegraf-nightly_freebsd_i386.tar.gz)
- [telegraf_nightly_i386.deb](https://dl.influxdata.com/telegraf/nightlies/telegraf_nightly_i386.deb)
- [telegraf-nightly.i386.rpm](https://dl.influxdata.com/telegraf/nightlies/telegraf-nightly.i386.rpm)
- [telegraf-nightly_linux_amd64.tar.gz](https://dl.influxdata.com/telegraf/nightlies/telegraf-nightly_linux_amd64.tar.gz)
- [telegraf-nightly_linux_arm64.tar.gz](https://dl.influxdata.com/telegraf/nightlies/telegraf-nightly_linux_arm64.tar.gz)
- [telegraf-nightly_linux_armel.tar.gz](https://dl.influxdata.com/telegraf/nightlies/telegraf-nightly_linux_armel.tar.gz)
- [telegraf-nightly_linux_armhf.tar.gz](https://dl.influxdata.com/telegraf/nightlies/telegraf-nightly_linux_armhf.tar.gz)
- [telegraf-nightly_linux_i386.tar.gz](https://dl.influxdata.com/telegraf/nightlies/telegraf-nightly_linux_i386.tar.gz)
- [telegraf-nightly_linux_s390x.tar.gz](https://dl.influxdata.com/telegraf/nightlies/telegraf-nightly_linux_s390x.tar.gz)
- [telegraf_nightly_s390x.deb](https://dl.influxdata.com/telegraf/nightlies/telegraf_nightly_s390x.deb)
- [telegraf-nightly.s390x.rpm](https://dl.influxdata.com/telegraf/nightlies/telegraf-nightly.s390x.rpm)
- [telegraf-nightly_windows_amd64.zip](https://dl.influxdata.com/telegraf/nightlies/telegraf-nightly_windows_amd64.zip)
- [telegraf-nightly_windows_i386.zip](https://dl.influxdata.com/telegraf/nightlies/telegraf-nightly_windows_i386.zip)
- [telegraf-nightly.x86_64.rpm](https://dl.influxdata.com/telegraf/nightlies/telegraf-nightly.x86_64.rpm)
- [telegraf-static-nightly_linux_amd64.tar.gz](https://dl.influxdata.com/telegraf/nightlies/telegraf-static-nightly_linux_amd64.tar.gz)

## How to use it:

See usage with:

```
telegraf --help
```

#### Generate a telegraf config file:

```
telegraf config > telegraf.conf
```

#### Generate config with only cpu input & influxdb output plugins defined:

```
telegraf --section-filter agent:inputs:outputs --input-filter cpu --output-filter influxdb config
```

#### Run a single telegraf collection, outputting metrics to stdout:

```
telegraf --config telegraf.conf --test
```

#### Run telegraf with all plugins defined in config file:

```
telegraf --config telegraf.conf
```

#### Run telegraf, enabling the cpu & memory input, and influxdb output plugins:

```
telegraf --config telegraf.conf --input-filter cpu:mem --output-filter influxdb
```

## Documentation

[Latest Release Documentation][release docs].

For documentation on the latest development code see the [documentation index][devel docs].

[release docs]: https://docs.influxdata.com/telegraf
[devel docs]: docs

## Input Plugins

* [sonic_telemery](./plugins/inputs/sonic_telemetry)

## Parsers

- [InfluxDB Line Protocol](/plugins/parsers/influx)
- [Collectd](/plugins/parsers/collectd)
- [CSV](/plugins/parsers/csv)
- [Dropwizard](/plugins/parsers/dropwizard)
- [FormUrlencoded](/plugins/parser/form_urlencoded)
- [Graphite](/plugins/parsers/graphite)
- [Grok](/plugins/parsers/grok)
- [JSON](/plugins/parsers/json)
- [Logfmt](/plugins/parsers/logfmt)
- [Nagios](/plugins/parsers/nagios)
- [Value](/plugins/parsers/value), ie: 45 or "booyah"
- [Wavefront](/plugins/parsers/wavefront)

## Serializers

- [InfluxDB Line Protocol](/plugins/serializers/influx)
- [JSON](/plugins/serializers/json)
- [Graphite](/plugins/serializers/graphite)
- [ServiceNow](/plugins/serializers/nowmetric)
- [SplunkMetric](/plugins/serializers/splunkmetric)
- [Carbon2](/plugins/serializers/carbon2)
- [Wavefront](/plugins/serializers/wavefront)

## Processor Plugins

* [clone](/plugins/processors/clone)
* [converter](/plugins/processors/converter)
* [date](/plugins/processors/date)
* [dedup](/plugins/processors/dedup)
* [defaults](/plugins/processors/defaults)
* [enum](/plugins/processors/enum)
* [execd](/plugins/processors/execd)
* [ifname](/plugins/processors/ifname)
* [filepath](/plugins/processors/filepath)
* [override](/plugins/processors/override)
* [parser](/plugins/processors/parser)
* [pivot](/plugins/processors/pivot)
* [port_name](/plugins/processors/port_name)
* [printer](/plugins/processors/printer)
* [regex](/plugins/processors/regex)
* [rename](/plugins/processors/rename)
* [reverse_dns](/plugins/processors/reverse_dns)
* [s2geo](/plugins/processors/s2geo)
* [starlark](/plugins/processors/starlark)
* [strings](/plugins/processors/strings)
* [tag_limit](/plugins/processors/tag_limit)
* [template](/plugins/processors/template)
* [topk](/plugins/processors/topk)
* [unpivot](/plugins/processors/unpivot)

## Aggregator Plugins

* [basicstats](./plugins/aggregators/basicstats)
* [final](./plugins/aggregators/final)
* [histogram](./plugins/aggregators/histogram)
* [merge](./plugins/aggregators/merge)
* [minmax](./plugins/aggregators/minmax)
* [valuecounter](./plugins/aggregators/valuecounter)

## Output Plugins

* [influxdb](./plugins/outputs/influxdb) (InfluxDB 1.x)
* [influxdb_v2](./plugins/outputs/influxdb_v2) ([InfluxDB 2.x](https://github.com/influxdata/influxdb))
* [amon](./plugins/outputs/amon)
* [amqp](./plugins/outputs/amqp) (rabbitmq)
* [application_insights](./plugins/outputs/application_insights)
* [aws kinesis](./plugins/outputs/kinesis)
* [aws cloudwatch](./plugins/outputs/cloudwatch)
* [azure_monitor](./plugins/outputs/azure_monitor)
* [cloud_pubsub](./plugins/outputs/cloud_pubsub) Google Cloud Pub/Sub
* [cratedb](./plugins/outputs/cratedb)
* [datadog](./plugins/outputs/datadog)
* [discard](./plugins/outputs/discard)
* [dynatrace](./plugins/outputs/dynatrace)
* [elasticsearch](./plugins/outputs/elasticsearch)
* [exec](./plugins/outputs/exec)
* [execd](./plugins/outputs/execd)
* [file](./plugins/outputs/file)
* [graphite](./plugins/outputs/graphite)
* [graylog](./plugins/outputs/graylog)
* [health](./plugins/outputs/health)
* [http](./plugins/outputs/http)
* [instrumental](./plugins/outputs/instrumental)
* [kafka](./plugins/outputs/kafka)
* [librato](./plugins/outputs/librato)
* [logz.io](./plugins/outputs/logzio)
* [mqtt](./plugins/outputs/mqtt)
* [nats](./plugins/outputs/nats)
* [newrelic](./plugins/outputs/newrelic)
* [nsq](./plugins/outputs/nsq)
* [opentsdb](./plugins/outputs/opentsdb)
* [prometheus](./plugins/outputs/prometheus_client)
* [riemann](./plugins/outputs/riemann)
* [riemann_legacy](./plugins/outputs/riemann_legacy)
* [socket_writer](./plugins/outputs/socket_writer)
* [stackdriver](./plugins/outputs/stackdriver) (Google Cloud Monitoring)
* [syslog](./plugins/outputs/syslog)
* [tcp](./plugins/outputs/socket_writer)
* [udp](./plugins/outputs/socket_writer)
* [warp10](./plugins/outputs/warp10)
* [wavefront](./plugins/outputs/wavefront)
* [sumologic](./plugins/outputs/sumologic)
* [yandex_cloud_monitoring](./plugins/outputs/yandex_cloud_monitoring)
