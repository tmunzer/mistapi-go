
# Webhook Location Asset

Location Asset sample

## Structure

`WebhookLocationAsset`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookLocationAssetEvent`](../../doc/models/webhook-location-asset-event.md) | Required | list of events |
| `Topic` | `string` | Required | topic subscribed to<br>**Default**: `"location_asset"` |

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
      "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
      "timestamp": 1461220784,
      "type": "asset",
      "x": 13.5,
      "y": 3.2
    }
  ],
  "topic": "location_asset"
}
```

