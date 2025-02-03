
# Installer Provision Device

Provision Device

*This model accepts additional fields of type interface{}.*

## Structure

`InstallerProvisionDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceprofileName` | `*string` | Optional | - |
| `ForSite` | `*bool` | Optional | - |
| `Height` | `*float64` | Optional | - |
| `MapId` | `*uuid.UUID` | Optional | - |
| `Name` | `string` | Required | - |
| `Orientation` | `*int` | Optional | - |
| `ReplacingMac` | `*string` | Optional | Onlif this is to replace an existing device |
| `Role` | `*string` | Optional | Optional role for switch / gateway |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SiteName` | `*string` | Optional | - |
| `X` | `*float64` | Optional | - |
| `Y` | `*float64` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "deviceprofile_name": "SJ1",
  "height": 2.7,
  "map_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
  "name": "SJ1-AP1",
  "orientation": 90,
  "replacing_mac": "5c5b3500003",
  "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
  "site_name": "SJ1",
  "x": 150,
  "y": 300,
  "for_site": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

