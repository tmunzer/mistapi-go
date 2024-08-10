
# Rssi Zone

RSSI Zone

## Structure

`RssiZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `Devices` | [`[]models.RssiZoneDevice`](../../doc/models/rssi-zone-device.md) | Required | List of devices and the respective RSSI values to be considered in the zone<br>**Constraints**: *Unique Items Required* |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `*string` | Optional | The name of the zone |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "devices": [
    {
      "device_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "rssi": 0
    }
  ],
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "name": "zone name",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "created_time": 116.44,
  "for_site": false,
  "modified_time": 218.52
}
```

