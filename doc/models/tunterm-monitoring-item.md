
# Tunterm Monitoring Item

## Structure

`TuntermMonitoringItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | can be ip, ipv6, hostname<br>**Constraints**: *Minimum Length*: `1` |
| `Port` | `*int` | Optional | when `protocol`==`tcp` |
| `Protocol` | [`*models.TunternMonitoringProtocolEnum`](../../doc/models/tuntern-monitoring-protocol-enum.md) | Optional | enum: `arp`, `ping`, `tcp`<br>**Constraints**: *Minimum Length*: `1` |
| `SrcVlanId` | `*int` | Optional | optional source for the monitoring check, vlan_id configured in tunterm_other_ip_configs |
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

