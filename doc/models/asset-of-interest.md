
# Asset of Interest

## Structure

`AssetOfInterest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Beam` | `*float64` | Optional | - |
| `By` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `CurrSite` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `DeviceName` | `*string` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `LastSeen` | `models.Optional[float64]` | Optional | Last seen timestamp |
| `Mac` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Manufacture` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `MapId` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Name` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Rssi` | `*float64` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "last_seen": 1470417522,
  "ap_mac": "ap_mac4",
  "beam": 115.36,
  "by": "by4",
  "curr_site": "curr_site8",
  "device_name": "device_name0"
}
```

