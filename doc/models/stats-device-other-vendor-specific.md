
# Stats Device Other Vendor Specific

When `vendor`==`cradlepoint`, contains Cradlepoint-specific statistics reported for the other device

## Structure

`StatsDeviceOtherVendorSpecific`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Interfaces` | [`map[string]models.StatsDeviceOtherVendorSpecificPort`](../../doc/models/stats-device-other-vendor-specific-port.md) | Optional | Cradlepoint interface statistics keyed by vendor interface identifier |
| `TargetVersion` | `*string` | Optional | Cradlepoint software version targeted for the device |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsDeviceOtherVendorSpecific := models.StatsDeviceOtherVendorSpecific{
        Interfaces:           map[string]models.StatsDeviceOtherVendorSpecificPort{
            "mdm-4d0e073b": models.StatsDeviceOtherVendorSpecificPort{
                BytesIn:              models.ToPointer(int64(5623096929)),
                BytesOut:             models.ToPointer(int64(12372750366)),
                Carrier:              models.ToPointer("Orange"),
                DisplayName:          models.ToPointer("display_name4"),
                Imei:                 models.ToPointer("866401234567893"),
                Imsi:                 models.ToPointer("2080101234567893"),
                Ip:                   models.ToPointer("10.134.237.57"),
                Link:                 models.ToPointer(true),
                Mode:                 models.ToPointer("wan"),
                Rsrp:                 models.ToPointer(float64(-108)),
                Rsrq:                 models.ToPointer(float64(-14)),
                Rssi:                 models.ToPointer(-74),
                ServiceMode:          models.ToPointer("5G NSA"),
                Sinr:                 models.ToPointer(float64(-1.2)),
                State:                models.ToPointer("READY"),
                Type:                 models.ToPointer("mdm"),
                Uptime:               models.ToPointer(2095779),
            },
        },
        TargetVersion:        models.ToPointer("7.23.40"),
    }

}
```

