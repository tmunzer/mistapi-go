
# Response Map Import

## Structure

`ResponseMapImport`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aps` | [`[]models.ResponseMapImportAp`](../../doc/models/response-map-import-ap.md) | Required | **Constraints**: *Unique Items Required* |
| `Floorplans` | [`[]models.ResponseMapImportFloorplan`](../../doc/models/response-map-import-floorplan.md) | Required | **Constraints**: *Unique Items Required* |
| `ForSite` | `*bool` | Optional | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `Summary` | [`models.ResponseMapImportSummary`](../../doc/models/response-map-import-summary.md) | Required | - |

## Example (as JSON)

```json
{
  "aps": [
    {
      "action": "named-placed",
      "floorplan_id": "000013fe-0000-0000-0000-000000000000",
      "height": 205.3,
      "mac": "mac8",
      "map_id": "00000c12-0000-0000-0000-000000000000",
      "orientation": 126,
      "reason": "reason0"
    }
  ],
  "floorplans": [
    {
      "action": "action6",
      "id": "000021bc-0000-0000-0000-000000000000",
      "map_id": "00002380-0000-0000-0000-000000000000",
      "name": "name6",
      "reason": "reason2"
    }
  ],
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "summary": {
    "num_ap_assigned": 66,
    "num_inv_assigned": 174,
    "num_map_assigned": 164
  },
  "for_site": false
}
```

