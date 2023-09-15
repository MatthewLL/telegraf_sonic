# SONiC GNMI Telemetry

SONiC GNMI Telemetry is an input plugin that consumes telemetry data based on the [GNMI](https://github.com/openconfig/reference/blob/master/rpc/gnmi/gnmi-specification.md) Subscribe method. TLS is supported for authentication and encryption.

It has been optimized to support GNMI telemetry as produced by SONiC.


### Configuration
SONiC---Telegraf---InfluxDB
```toml
[[inputs.sonic_telemetry_gnmi]]
  ## Address and port of the GNMI GRPC server，which should be a SONiC Switch with Telemetry enabled，port is 8080 as default
  addresses = ["172.20.222.97:8085"]

  ## define credentials
  #username = "admin"
  #password = "YourPaSsWoRd"

  ## GNMI encoding requested (one of: "proto", "json", "json_ietf")
  encoding = "json_ietf"

  ## redial in case of failures after
  redial = "10s"

  ## enable client-side TLS and define CA to authenticate the device
  #enable_tls = true
  # tls_ca = "/etc/telegraf/ca.pem"
  #insecure_skip_verify = true

  ## define client-side TLS certificate & key to authenticate to the device
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"

  ## GNMI subscription prefix (optional, can usually be left empty)
  ## See: https://github.com/openconfig/reference/blob/master/rpc/gnmi/gnmi-specification.md#222-paths
  # origin = ""
  # prefix = ""
  #target is a database name or "OTHERS"
  target = "COUNTERS_DB"

  ## Define additional aliases to map telemetry encoding paths to simple measurement names
  # [inputs.sonic_telemetry_gnmi.aliases]
  #   ifcounters = "openconfig:/interfaces/interface/state/counters"

  [[inputs.sonic_telemetry_gnmi.subscription]]
    ## Name of the measurement that will be emitted
    name = "ifcounters"

    ## Origin and path of the subscription
    ## See: https://github.com/openconfig/reference/blob/master/rpc/gnmi/gnmi-specification.md#222-paths
    ##
    ## origin usually refers to a (YANG) data model implemented by the device
    ## and path to a specific substructe inside it that should be subscribed to (similar to an XPath)
    ## YANG models can be found e.g. here: https://sonic_mgmt_ip_address/ui
    origin = ""
    path = "/COUNTERS/Ethernet0"

    # Subscription mode (one of: "target_defined", "sample", "on_change") and interval
    subscription_mode = "sample"
    sample_interval = "5s"

    ## Suppress redundant transmissions when measured values are unchanged
    # suppress_redundant = false

    ## If suppression is enabled, send updates at least every X seconds anyway
    # heartbeat_interval = "60s"

[[outputs.influxdb_v2]]
  ## The URLs of the InfluxDB cluster nodes.
  ##
  ## Multiple URLs can be specified for a single cluster, only ONE of the
  ## urls will be written to each interval.
  ##   ex: urls = ["https://us-west-2-1.aws.cloud2.influxdata.com"]
  urls = ["http://10.69.65.19:8086"]

  ## Token for authentication.
  token = "E6u8WrhRafHZdPQ1JUyQY3C1DQKe37crtlPlE613oq8voyeoRrP3NLy7FGRFt84XEtHIvlRpIxZ4EB-Qns2C7w=="

  ## Organization is the name of the organization you wish to write to; must exist.
  organization = "lmf-Organization"

  ## Destination bucket to write into.
  bucket = "test"

  ## The value of this tag will be used to determine the bucket.  If this
  ## tag is not set the 'bucket' option is used as the default.
  # bucket_tag = ""

  ## If true, the bucket tag will not be added to the metric.
  # exclude_bucket_tag = false

  ## Timeout for HTTP messages.
  # timeout = "5s"

  ## Additional HTTP headers
  # http_headers = {"X-Special-Header" = "Special-Value"}

  ## HTTP Proxy override, if unset values the standard proxy environment
  ## variables are consulted to determine which proxy, if any, should be used.
  # http_proxy = "http://corporate.proxy:3128"

  ## HTTP User-Agent
  # user_agent = "telegraf"

  ## Content-Encoding for write request body, can be set to "gzip" to
  ## compress body or "identity" to apply no encoding.
  # content_encoding = "gzip"

  ## Enable or disable uint support for writing uints influxdb 2.0.
  # influx_uint_support = false

  ## Optional TLS Config for use on HTTP connections.
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false
```
### Validation
You should be able to find measurement "ifcounters" in database "sonic_telemetry",which could be treated as Grafana's data source.



