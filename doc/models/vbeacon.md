
# Vbeacon

vBeacon

## Structure

`Vbeacon`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `Major` | `*int` | Optional | bluetooth tag major |
| `MapId` | `*uuid.UUID` | Optional | map where the device belongs to |
| `Message` | `*string` | Optional | a message that can be displayed when the sdkclient gets near the vbeacon |
| `Minor` | `*int` | Optional | bluetooth tag minor |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `*string` | Optional | name / label of the device |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Power` | `*int` | Optional | required if `power_mode`==`custom`, -30 - 100, in dBm. For default power_mode, power = 4 dBm.<br>**Default**: `4`<br>**Constraints**: `>= -30`, `<= 100` |
| `PowerMode` | [`*models.BleConfigPowerModeEnum`](../../doc/models/ble-config-power-mode-enum.md) | Optional | enum: `custom`, `default`<br>**Default**: `"default"` |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Url` | `*string` | Optional | URL to show, optional |
| `Uuid` | `*uuid.UUID` | Optional | bluetooth tag UUID |
| `WayfindingNodename` | `*string` | Optional | the name to be used in wayfinding_path or wayfinding_grid blob |
| `X` | `*float64` | Optional | x in pixel |
| `Y` | `*float64` | Optional | y in pixel |

## Example (as JSON)

```json
{
  "major": 1356,
  "map_id": "63eda950-c6da-11e4-a628-60f81dd250cc",
  "message": "Welcome to Mist",
  "minor": 21,
  "name": "conference room",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "power": 4,
  "power_mode": "custom",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "url": "http://www.mist.com/any",
  "uuid": "31375aeb-b8d3-1ea6-83bf-a31eb04e1c38",
  "wayfinding_nodename": "node1",
  "x": 53.5,
  "y": 173.1,
  "created_time": 135.22,
  "for_site": false,
  "id": "000014ac-0000-0000-0000-000000000000"
}
```

