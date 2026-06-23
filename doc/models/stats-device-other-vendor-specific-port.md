
# Stats Device Other Vendor Specific Port

Cradlepoint interface statistics for a device tracked by Mist

## Structure

`StatsDeviceOtherVendorSpecificPort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BytesIn` | `*int64` | Optional | Total bytes received on the Cradlepoint interface |
| `BytesOut` | `*int64` | Optional | Total bytes transmitted from the Cradlepoint interface |
| `Carrier` | `*string` | Optional | Cellular carrier name reported for the Cradlepoint interface, when applicable |
| `DisplayName` | `*string` | Optional | Human-readable interface name reported by Cradlepoint |
| `Imei` | `*string` | Optional | Modem IMEI serving the Cradlepoint interface, when applicable |
| `Imsi` | `*string` | Optional | Subscriber IMSI associated with the Cradlepoint interface, when applicable |
| `Ip` | `*string` | Optional | Cradlepoint interface IP address reported by the vendor |
| `Link` | `*bool` | Optional | Whether the Cradlepoint interface link is up |
| `Mode` | `*string` | Optional | Interface role reported by Cradlepoint, such as wan or lan |
| `Mtu` | `*int` | Optional | Maximum transmission unit configured on the Cradlepoint interface, in bytes |
| `PortParent` | `*string` | Optional | Parent Cradlepoint port or interface group for this interface |
| `Rsrp` | `*float64` | Optional | Reference signal received power for a Cradlepoint cellular interface, in dBm |
| `Rsrq` | `*float64` | Optional | Reference signal received quality for a Cradlepoint cellular interface, in dB |
| `Rssi` | `*int` | Optional | Received signal strength indicator for a Cradlepoint cellular interface, in dBm |
| `ServiceMode` | `*string` | Optional | Cellular service mode reported for the Cradlepoint interface, such as 5G NSA |
| `Sinr` | `*float64` | Optional | Signal-to-interference-plus-noise ratio for a Cradlepoint cellular interface, in dB |
| `State` | `*string` | Optional | Operational state reported for the Cradlepoint interface |
| `Type` | `*string` | Optional | Interface type reported by Cradlepoint, such as mdm or ethernet |
| `Uptime` | `*int` | Optional | Elapsed time since the Cradlepoint interface last became active, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsDeviceOtherVendorSpecificPort := models.StatsDeviceOtherVendorSpecificPort{
        BytesIn:              models.ToPointer(int64(5623096929)),
        BytesOut:             models.ToPointer(int64(12372750366)),
        Carrier:              models.ToPointer("Orange"),
        DisplayName:          models.ToPointer("mdm-4d0e073b"),
        Imei:                 models.ToPointer("866401234567893"),
        Imsi:                 models.ToPointer("2080101234567893"),
        Ip:                   models.ToPointer("10.134.237.57"),
        Link:                 models.ToPointer(true),
        Mode:                 models.ToPointer("wan"),
        Mtu:                  models.ToPointer(1500),
        PortParent:           models.ToPointer("mdm"),
        Rsrp:                 models.ToPointer(float64(-108)),
        Rsrq:                 models.ToPointer(float64(-14)),
        Rssi:                 models.ToPointer(-74),
        ServiceMode:          models.ToPointer("5G NSA"),
        Sinr:                 models.ToPointer(float64(-1.2)),
        State:                models.ToPointer("READY"),
        Type:                 models.ToPointer("mdm"),
        Uptime:               models.ToPointer(2095779),
    }

}
```

