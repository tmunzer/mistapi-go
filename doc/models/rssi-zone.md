
# Rssi Zone

RSSI Zone

*This model accepts additional fields of type interface{}.*

## Structure

`RssiZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `Devices` | [`[]models.RssiZoneDevice`](../../doc/models/rssi-zone-device.md) | Required | List of devices and the respective RSSI values to be considered in the zone<br>**Constraints**: *Unique Items Required* |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | The name of the zone |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "devices": [
    {
      "device_id": "00000000-0000-0000-1000-d8695a0f9e61",
      "rssi": 0,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "zone name",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "created_time": 116.44,
  "for_site": false,
  "modified_time": 218.52,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

