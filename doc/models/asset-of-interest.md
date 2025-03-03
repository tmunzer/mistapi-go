
# Asset of Interest

*This model accepts additional fields of type interface{}.*

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
| `LastSeen` | `*float64` | Optional | - |
| `Mac` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Manufacture` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `MapId` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Name` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Rssi` | `*float64` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "ap_mac": "ap_mac4",
  "beam": 115.36,
  "by": "by4",
  "curr_site": "curr_site8",
  "device_name": "device_name0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

