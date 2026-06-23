
# Stats Device Other

Statistics for a third-party or other device tracked by Mist

## Structure

`StatsDeviceOther`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CachedStats` | `*bool` | Optional | Whether the response was served from cached vendor statistics |
| `ConfigStatus` | `*string` | Optional | Synchronization status of the other device configuration in Mist |
| `ConnectedDevices` | [`map[string]models.StatsDeviceOtherConnectedDevice`](../../doc/models/stats-device-other-connected-device.md) | Optional | Property key is the connected device MAC address |
| `Interfaces` | [`map[string]models.StatsDeviceOtherInterface`](../../doc/models/stats-device-other-interface.md) | Optional | Property key is the interface name |
| `LastConfig` | `*int` | Optional | Timestamp of the last configuration update reported for the other device, in epoch seconds |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `LldpEnabled` | `*bool` | Optional | Whether LLDP is enabled on the other device |
| `Mac` | `*string` | Optional | Other device MAC address reported in the statistics response |
| `Status` | `*string` | Optional | Connectivity status reported for the other device, such as online |
| `Uptime` | `*int` | Optional | Elapsed time since the other device last booted, in seconds |
| `Vendor` | `*string` | Optional | Device vendor name, such as Cradlepoint |
| `VendorSpecific` | [`*models.StatsDeviceOtherVendorSpecific`](../../doc/models/stats-device-other-vendor-specific.md) | Optional | When `vendor`==`cradlepoint`, contains Cradlepoint-specific statistics reported for the other device |
| `Version` | `*string` | Optional | Vendor software version reported for the other device |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsDeviceOther := models.StatsDeviceOther{
        CachedStats:          models.ToPointer(false),
        ConfigStatus:         models.ToPointer("synced"),
        ConnectedDevices:     map[string]models.StatsDeviceOtherConnectedDevice{
            "key0": models.StatsDeviceOtherConnectedDevice{
                Mac:                  models.ToPointer("mac2"),
                Name:                 models.ToPointer("name8"),
                PortId:               models.ToPointer("port_id8"),
                Type:                 models.ToPointer("type8"),
            },
            "key1": models.StatsDeviceOtherConnectedDevice{
                Mac:                  models.ToPointer("mac2"),
                Name:                 models.ToPointer("name8"),
                PortId:               models.ToPointer("port_id8"),
                Type:                 models.ToPointer("type8"),
            },
            "key2": models.StatsDeviceOtherConnectedDevice{
                Mac:                  models.ToPointer("mac2"),
                Name:                 models.ToPointer("name8"),
                PortId:               models.ToPointer("port_id8"),
                Type:                 models.ToPointer("type8"),
            },
        },
        Interfaces:           map[string]models.StatsDeviceOtherInterface{
            "key0": models.StatsDeviceOtherInterface{
                BytesIn:              models.ToPointer(int64(48)),
                BytesOut:             models.ToPointer(int64(188)),
                Carrier:              models.ToPointer("carrier6"),
                Imei:                 models.ToPointer("imei8"),
                Imsi:                 models.ToPointer("imsi2"),
            },
            "key1": models.StatsDeviceOtherInterface{
                BytesIn:              models.ToPointer(int64(48)),
                BytesOut:             models.ToPointer(int64(188)),
                Carrier:              models.ToPointer("carrier6"),
                Imei:                 models.ToPointer("imei8"),
                Imsi:                 models.ToPointer("imsi2"),
            },
            "key2": models.StatsDeviceOtherInterface{
                BytesIn:              models.ToPointer(int64(48)),
                BytesOut:             models.ToPointer(int64(188)),
                Carrier:              models.ToPointer("carrier6"),
                Imei:                 models.ToPointer("imei8"),
                Imsi:                 models.ToPointer("imsi2"),
            },
        },
        LastConfig:           models.ToPointer(1675392788),
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        Mac:                  models.ToPointer("5c5b35000018"),
        Status:               models.ToPointer("online"),
        Uptime:               models.ToPointer(20296),
        Vendor:               models.ToPointer("cradlepoint"),
        Version:              models.ToPointer("7.22.70"),
    }

}
```

