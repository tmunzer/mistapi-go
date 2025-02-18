
# Webhook Location Asset

Location Asset sample

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookLocationAsset`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationAssetEvent`](../../doc/models/webhook-location-asset-event.md) | Required | List of events |
| `Topic` | `string` | Required | Topic subscribed to<br>**Default**: `"location_asset"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "events": [
    {
      "battery_voltage": 3370,
      "eddystone_uid_instance": "5c5b35000001",
      "eddystone_uid_namespace": "2818e3868dec25629ede",
      "eddystone_url_url": "https://www.abc.com",
      "ibeacon_major": 13,
      "ibeacon_minor": 138,
      "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
      "mac": "7fc2936fd243",
      "map_id": "845a23bf-bed9-e43c-4c86-6fa474be7ae5",
      "mfg_company_id": 935,
      "mfg_data": "648520a1020000",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "timestamp": 1461220784,
      "type": "asset",
      "x": 13.5,
      "y": 3.2,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "topic": "location_asset",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

