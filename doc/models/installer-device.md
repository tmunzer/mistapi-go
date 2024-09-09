
# Installer Device

## Structure

`InstallerDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Connected` | `*bool` | Optional | - |
| `DeviceprofileName` | `*string` | Optional | - |
| `ExtIp` | `*string` | Optional | - |
| `Height` | `*float64` | Optional | - |
| `Ip` | `*string` | Optional | - |
| `LastSeen` | `*float64` | Optional | - |
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
  "last_seen": 1687887907.59198,
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
  "y": 300
}
```

