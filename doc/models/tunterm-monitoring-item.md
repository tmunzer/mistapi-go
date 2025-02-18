
# Tunterm Monitoring Item

*This model accepts additional fields of type interface{}.*

## Structure

`TuntermMonitoringItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | Can be ip, ipv6, hostname<br>**Constraints**: *Minimum Length*: `1` |
| `Port` | `*int` | Optional | When `protocol`==`tcp` |
| `Protocol` | [`*models.TunternMonitoringProtocolEnum`](../../doc/models/tuntern-monitoring-protocol-enum.md) | Optional | enum: `arp`, `ping`, `tcp`<br>**Constraints**: *Minimum Length*: `1` |
| `SrcVlanId` | `*int` | Optional | Optional source for the monitoring check, vlan_id configured in tunterm_other_ip_configs |
| `Timeout` | `*int` | Optional | **Default**: `300` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "host": "10.2.8.15",
  "port": 80,
  "protocol": "tcp",
  "src_vlan_id": 5,
  "timeout": 300,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

