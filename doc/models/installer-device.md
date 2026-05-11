
# Installer Device

## Structure

`InstallerDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BleStat` | [`*models.InstallerDeviceBleStat`](../../doc/models/installer-device-ble-stat.md) | Optional | BLE statistics for the device |
| `Connected` | `*bool` | Optional | - |
| `DeviceprofileName` | `*string` | Optional | - |
| `ExtIp` | `*string` | Optional | - |
| `Height` | `*float64` | Optional | - |
| `Ip` | `*string` | Optional | - |
| `LastSeen` | `models.Optional[float64]` | Optional | Last seen timestamp |
| `Mac` | `*string` | Optional | - |
| `MapId` | `*uuid.UUID` | Optional | - |
| `Model` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Orientation` | `*int` | Optional | - |
| `Serial` | `*string` | Optional | - |
| `SiteName` | `*string` | Optional | - |
| `Uptime` | `*int` | Optional | - |
| `VcMac` | `models.Optional[string]` | Optional | - |
| `Version` | `*string` | Optional | - |
| `X` | `*float64` | Optional | - |
| `Y` | `*float64` | Optional | - |

## Example (as JSON)

```json
{
  "connected": true,
  "deviceprofile_name": "SJ1",
  "ext_ip": "12.34.56.78",
  "height": 2.7,
  "ip": "192.168.1.111",
  "last_seen": 1470417522,
  "mac": "5c5b35000018",
  "map_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
  "model": "AP41",
  "name": "hallway",
  "orientation": 90,
  "serial": "FXLH2015150025",
  "site_name": "SJ1",
  "uptime": 12345,
  "version": "0.10.24362",
  "x": 150,
  "y": 300,
  "ble_stat": {
    "major": 210,
    "minors": [
      237,
      238
    ],
    "uuid": "00001d0e-0000-0000-0000-000000000000"
  }
}
```

