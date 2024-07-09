
# Tunterm Monitoring Item

## Structure

`TuntermMonitoringItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | can be ip, ipv6, hostname<br>**Constraints**: *Minimum Length*: `1` |
| `Port` | `*int` | Optional | when `protocol`==`tcp` |
| `Protocol` | [`*models.TunternMonitoringProtocolEnum`](../../doc/models/tuntern-monitoring-protocol-enum.md) | Optional | **Constraints**: *Minimum Length*: `1` |
| `Timeout` | `*int` | Optional | **Default**: `300` |

## Example (as JSON)

```json
{
  "host": "10.2.8.15",
  "port": 80,
  "protocol": "tcp",
  "timeout": 300
}
```

