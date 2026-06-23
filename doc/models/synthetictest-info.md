
# Synthetictest Info

Synthetic test status or result record

## Structure

`SynthetictestInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `By` | `*string` | Optional | Actor that triggered the synthetic test |
| `DeviceType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Optional | enum: `ap`, `gateway`, `switch` |
| `Failed` | `*bool` | Optional | Whether the synthetic test failed |
| `Latency` | `*int` | Optional | Measured latency for the synthetic test, in milliseconds |
| `Mac` | `*string` | Optional | Device MAC address that ran the synthetic test |
| `PortId` | `*string` | Optional | Interface identifier used to run the synthetic test |
| `Reason` | `*string` | Optional | Failure reason reported for the synthetic test, when available |
| `RxMbps` | `*int` | Optional | Receive throughput measured by the synthetic speed test, in Mbps |
| `StartTime` | `*int` | Optional | Epoch timestamp, in seconds, when the synthetic test started |
| `Status` | `*string` | Optional | Current execution status of the synthetic test |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `TxMbps` | `*int` | Optional | Transmit throughput measured by the synthetic speed test, in Mbps |
| `Type` | [`*models.SynthetictestTypeEnum`](../../doc/models/synthetictest-type-enum.md) | Optional | enum: `arp`, `curl`, `dhcp`, `dhcp6`, `dns`, `lan_connectivity`, `radius`, `speedtest` |
| `VlanId` | `*int` | Optional | Identifier of the VLAN used by the synthetic test |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    synthetictestInfo := models.SynthetictestInfo{
        By:                   models.ToPointer("user"),
        DeviceType:           models.ToPointer(models.DeviceTypeEnum_AP),
        Failed:               models.ToPointer(false),
        Latency:              models.ToPointer(40),
        Mac:                  models.ToPointer("mac2"),
        PortId:               models.ToPointer("ge-0/0/2"),
        Reason:               models.ToPointer("interface not ready to perform test"),
        RxMbps:               models.ToPointer(322),
        StartTime:            models.ToPointer(1675718807),
        TxMbps:               models.ToPointer(199),
        VlanId:               models.ToPointer(20),
    }

}
```

