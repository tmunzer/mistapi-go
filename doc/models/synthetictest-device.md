
# Synthetictest Device

*This model accepts additional fields of type interface{}.*

## Structure

`SynthetictestDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | If `type`==`lan_connectivity` |
| `Hostname` | `*string` | Optional | If `type`==`dns` |
| `Ip` | `*string` | Optional | If `type`==`arp` |
| `Password` | `*string` | Optional | If `type`==`radius` |
| `PingCount` | `*int` | Optional | If `type`==`lan_connectivity`<br>**Default**: `10`<br>**Constraints**: `>= 10`, `<= 60` |
| `PingDetails` | `*bool` | Optional | If `type`==`lan_connectivity`<br>**Default**: `false` |
| `PingSize` | `*int` | Optional | If `type`==`lan_connectivity`<br>**Default**: `56`<br>**Constraints**: `>= 56`, `<= 65535` |
| `PortId` | `*string` | Optional | If `type`==`speedtest`, required for ssr |
| `Protocol` | [`*models.SynthetictestDeviceProtocolEnum`](../../doc/models/synthetictest-device-protocol-enum.md) | Optional | if `type`==`lan_connectivity`. enum: `ping`, `traceroute`, `ping+traceroute`<br>**Default**: `"ping+traceroute"` |
| `Tenant` | `*string` | Optional | If `type`==`lan_connectivity` |
| `TracerouteUdpPort` | `*int` | Optional | SRX only, traceroute udp port<br>**Default**: `33434`<br>**Constraints**: `>= 0`, `<= 65535` |
| `Type` | [`models.SynthetictestTypeEnum`](../../doc/models/synthetictest-type-enum.md) | Required | enum: `arp`, `curl`, `dhcp`, `dhcp6`, `dns`, `lan_connectivity`, `radius`, `speedtest` |
| `Url` | `*string` | Optional | If `type`==`curl` |
| `Username` | `*string` | Optional | If `type`==`radius` |
| `VlanId` | [`*models.SynthetictestDeviceVlanId`](../../doc/models/containers/synthetictest-device-vlan-id.md) | Optional | Required for AP |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "host": "www.example.com",
  "hostname": "google.com\"",
  "ip": "192.168.3.5",
  "password": "test123",
  "ping_count": 10,
  "ping_details": false,
  "ping_size": 56,
  "port_id": "wan0",
  "protocol": "ping+traceroute",
  "tenant": "lan_network1",
  "traceroute_udp_port": 33434,
  "type": "dns",
  "url": "https://www.example.com",
  "username": "user",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

