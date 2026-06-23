
# Webhook Minis Application Event

Marvis Minis application synthetic test result

## Structure

`WebhookMinisApplicationEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceMac` | `*string` | Optional | MAC address of the device |
| `Ip` | `*string` | Optional | Destination IP address used for the application test |
| `Latency` | `*int` | Optional | Application test latency measured in milliseconds |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `ProbeName` | `*string` | Optional | Name of the probe |
| `ProbeType` | `*string` | Optional | Probe category used for the Minis application test |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `SrcIp` | `*string` | Optional | Source IP address of the test |
| `Success` | `*bool` | Optional | Whether the test was successful |
| `TestType` | [`*models.SynthetictestConfigCustomProbeTypeEnum`](../../doc/models/synthetictest-config-custom-probe-type-enum.md) | Optional | enum: `application`, `curl`, `icmp`, `reachability`, `tcp`<br><br>**Default**: `"icmp"` |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Vlan` | `*int` | Optional | Network VLAN ID used for the application test |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    webhookMinisApplicationEvent := models.WebhookMinisApplicationEvent{
        DeviceMac:            models.ToPointer("device_mac0"),
        Ip:                   models.ToPointer("ip0"),
        Latency:              models.ToPointer(106),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        ProbeName:            models.ToPointer("connectivitycheck.gstatic.com"),
        ProbeType:            models.ToPointer("application"),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        TestType:             models.ToPointer(models.SynthetictestConfigCustomProbeTypeEnum_ICMP),
        Vlan:                 models.ToPointer(12),
    }

}
```

