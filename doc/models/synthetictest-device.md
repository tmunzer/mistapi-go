
# Synthetictest Device

Request body for triggering a synthetic test on a device

*This model accepts additional fields of type interface{}.*

## Structure

`SynthetictestDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | If `type`==`lan_connectivity`. Hostname or IP address to probe for LAN connectivity testing |
| `Hostname` | `*string` | Optional | If `type`==`dns`. Hostname to resolve during the DNS synthetic test |
| `Ip` | `*string` | Optional | If `type`==`arp`. IP address to resolve through ARP during the synthetic test |
| `Password` | `*string` | Optional | If `type`==`radius`. Password used for RADIUS authentication testing |
| `PingCount` | `*int` | Optional | If `type`==`lan_connectivity`. Number of ping probes to send during the LAN connectivity test<br><br>**Default**: `10`<br><br>**Constraints**: `>= 10`, `<= 60` |
| `PingDetails` | `*bool` | Optional | If `type`==`lan_connectivity`. Whether to include per-ping results in the LAN connectivity test output<br><br>**Default**: `false` |
| `PingSize` | `*int` | Optional | If `type`==`lan_connectivity`. Payload size, in bytes, for each ping probe<br><br>**Default**: `56`<br><br>**Constraints**: `>= 56`, `<= 65535` |
| `PortId` | `*string` | Optional | If `type`==`speedtest`, required for ssr |
| `Protocol` | [`*models.SynthetictestDeviceProtocolEnum`](../../doc/models/synthetictest-device-protocol-enum.md) | Optional | If `type`==`lan_connectivity`. Protocol or protocol combination used by the LAN connectivity test. enum: `ping`, `traceroute`, `ping+traceroute`<br><br>**Default**: `"ping+traceroute"` |
| `Tenant` | `*string` | Optional | If `type`==`curl` or `type`==`lan_connectivity` |
| `Timeout` | `*int` | Optional | If `type`==`curl`. Timeout, in seconds, for the HTTP request<br><br>**Default**: `60`<br><br>**Constraints**: `>= 30`, `<= 120` |
| `TracerouteUdpPort` | `*int` | Optional | SRX only, traceroute udp port<br><br>**Default**: `33434`<br><br>**Constraints**: `>= 0`, `<= 65535` |
| `Type` | [`models.SynthetictestTypeEnum`](../../doc/models/synthetictest-type-enum.md) | Required | enum: `arp`, `curl`, `dhcp`, `dhcp6`, `dns`, `lan_connectivity`, `radius`, `speedtest` |
| `Url` | `*string` | Optional | If `type`==`curl`. URL requested during the HTTP synthetic test |
| `Username` | `*string` | Optional | If `type`==`radius`. Username used for RADIUS authentication testing |
| `VlanId` | [`*models.SynthetictestDeviceVlanId`](../../doc/models/containers/synthetictest-device-vlan-id.md) | Optional | Required for AP. VLAN ID used by the synthetic test when the target device is an AP |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    synthetictestDevice := models.SynthetictestDevice{
        Host:                 models.ToPointer("www.example.com"),
        Hostname:             models.ToPointer("google.com\""),
        Ip:                   models.ToPointer("192.168.3.5"),
        Password:             models.ToPointer("test123"),
        PingCount:            models.ToPointer(10),
        PingDetails:          models.ToPointer(false),
        PingSize:             models.ToPointer(56),
        PortId:               models.ToPointer("wan0"),
        Protocol:             models.ToPointer(models.SynthetictestDeviceProtocolEnum_ENUMPINGTRACEROUTE),
        Tenant:               models.ToPointer("lan_network1"),
        Timeout:              models.ToPointer(60),
        TracerouteUdpPort:    models.ToPointer(33434),
        Type:                 models.SynthetictestTypeEnum_DNS,
        Url:                  models.ToPointer("https://www.example.com"),
        Username:             models.ToPointer("user"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

