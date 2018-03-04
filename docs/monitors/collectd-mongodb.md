<!--- GENERATED BY gomplate from scripts/docs/monitor-page.md.tmpl --->

# collectd/mongodb

 Monitors an instance of MongoDB using the
[collectd MongoDB Python plugin](https://github.com/signalfx/collectd-mongodb).

Also see https://github.com/signalfx/integrations/tree/master/collectd-mongodb.


Monitor Type: `collectd/mongodb`

[Monitor Source Code](https://github.com/signalfx/signalfx-agent/tree/master/internal/monitors/collectd/mongodb)

**Accepts Endpoints**: **Yes**

**Multiple Instances Allowed**: Yes

## Configuration

| Config option | Required | Type | Description |
| --- | --- | --- | --- |
| `host` | **yes** | `string` |  |
| `port` | **yes** | `integer` |  |
| `name` | no | `string` |  |
| `databases` | no | `list of string` |  |
| `username` | no | `string` |  |
| `password` | no | `string` |  |
| `useTLS` | no | `bool` |  (**default:** `false`) |
| `caCerts` | no | `string` |  |
| `tlsClientCert` | no | `string` |  |
| `tlsClientKey` | no | `string` |  |
| `tlsClientKeyPassPhrase` | no | `string` |  |





