
# Synthetictest Info

## Structure

`SynthetictestInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `By` | `*string` | Optional | - |
| `DeviceType` | [`*models.SynthetictestInfoDeviceTypeEnum`](../../doc/models/synthetictest-info-device-type-enum.md) | Optional | enum: `ap`, `gateway`, `switch` |
| `Failed` | `*bool` | Optional | - |
| `Latency` | `*int` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `PortId` | `*string` | Optional | - |
| `Reason` | `*string` | Optional | - |
| `RxMbps` | `*int` | Optional | - |
| `StartTime` | `*int` | Optional | - |
| `Status` | `*string` | Optional | - |
| `Timestamp` | `*float64` | Optional | - |
| `TxMbps` | `*int` | Optional | - |
| `Type` | [`*models.SynthetictestTypeEnum`](../../doc/models/synthetictest-type-enum.md) | Optional | enum: `arp`, `curl`, `dhcp`, `dhcp6`, `dns`, `radius`, `speedtest` |
| `VlanId` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "by": "user",
  "failed": false,
  "latency": 40,
  "port_id": "ge-0/0/2",
  "reason": "interface not ready to perform test",
  "rx_mbps": 322,
  "start_time": 1675718807,
  "timestamp": 1675718807,
  "tx_mbps": 199,
  "vlan_id": 20,
  "device_type": "switch",
  "mac": "mac2"
}
```

