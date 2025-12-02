
# Tunterm Monitoring Item

## Structure

`TuntermMonitoringItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | Can be ip, ipv6, hostname<br><br>**Constraints**: *Minimum Length*: `1` |
| `Port` | `*int` | Optional | When `protocol`==`tcp` |
| `Protocol` | [`*models.TuntermMonitoringProtocolEnum`](../../doc/models/tunterm-monitoring-protocol-enum.md) | Optional | enum: `arp`, `ping`, `tcp`<br><br>**Constraints**: *Minimum Length*: `1` |
| `SrcVlanId` | `*int` | Optional | Optional source for the monitoring check, vlan_id configured in tunterm_other_ip_configs |
| `Timeout` | `*int` | Optional | **Default**: `300` |

## Example (as JSON)

```json
{
  "host": "10.2.8.15",
  "port": 80,
  "protocol": "tcp",
  "src_vlan_id": 5,
  "timeout": 300
}
```

