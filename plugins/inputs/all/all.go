package all

import (
	
	_ "github.com/influxdata/telegraf/plugins/inputs/influxdb"
	_ "github.com/influxdata/telegraf/plugins/inputs/influxdb_listener"
	_ "github.com/influxdata/telegraf/plugins/inputs/influxdb_v2_listener"
	
	_ "github.com/influxdata/telegraf/plugins/inputs/sonic_telemetry"

)
