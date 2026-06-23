
# Stats Marvis Client

Marvis Client stats record returned by search

## Structure

`StatsMarvisClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BatteryCharging` | `*bool` | Optional | Whether the device battery is currently charging |
| `BatteryLevel` | `*int` | Optional | Battery level percentage (0–100) |
| `CpuBackground` | `*float64` | Optional | Background CPU utilization (0–100) |
| `CpuIdle` | `*float64` | Optional | Idle CPU percentage (0–100) |
| `CpuSystem` | `*float64` | Optional | System CPU utilization (0–100) |
| `CpuUser` | `*float64` | Optional | User-space CPU utilization (0–100) |
| `DeviceId` | `*uuid.UUID` | Optional | UUID of the device the Marvis Client is installed on |
| `Hostname` | `*string` | Optional | Device hostname |
| `Location` | [`*models.StatsMarvisClientLocation`](../../doc/models/stats-marvis-client-location.md) | Optional | Last known location fix for a Marvis Client device |
| `MemoryTotal` | `*int` | Optional | Total device memory, in bytes |
| `MemoryUsage` | `*int` | Optional | Memory in use, in bytes |
| `Mfg` | `*string` | Optional | Device manufacturer |
| `Model` | `*string` | Optional | Device model name |
| `OrgId` | `*uuid.UUID` | Optional | Organization UUID |
| `OsType` | `*string` | Optional | OS type or platform (e.g. Android, iOS) |
| `OsVersion` | `*string` | Optional | OS version string |
| `Serial` | `*string` | Optional | Device serial number |
| `StorageTotal` | `*int` | Optional | Total device storage, in bytes |
| `StorageUsage` | `*int` | Optional | Storage in use, in bytes |
| `Timestamp` | `*int` | Optional | Timestamp of the stats record, in epoch seconds |
| `WifiBand` | `*string` | Optional | Wi-Fi band the device is connected on |
| `WifiBssid` | `*string` | Optional | BSSID the device is connected to |
| `WifiChannel` | `*int` | Optional | Wi-Fi channel the device is on |
| `WifiIp` | `*string` | Optional | Device Wi-Fi IP address |
| `WifiMac` | `*string` | Optional | Device Wi-Fi MAC address |
| `WifiRssi` | `*int` | Optional | Wi-Fi RSSI, in dBm |
| `WifiSsid` | `*string` | Optional | SSID the device is connected to |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsMarvisClient := models.StatsMarvisClient{
        BatteryCharging:      models.ToPointer(false),
        BatteryLevel:         models.ToPointer(2),
        CpuBackground:        models.ToPointer(float64(228.2)),
        CpuIdle:              models.ToPointer(float64(142.1)),
        CpuSystem:            models.ToPointer(float64(19.82)),
    }

}
```

