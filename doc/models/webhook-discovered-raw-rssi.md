
# Webhook Discovered Raw Rssi

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookDiscoveredRawRssi`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookDiscoveredRawRssiEvent`](../../doc/models/webhook-discovered-raw-rssi-event.md) | Optional | - |
| `Topic` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "events": [
    {
      "ap_loc": [
        26.98,
        26.97,
        26.96
      ],
      "beam": 124,
      "device_id": "0000254a-0000-0000-0000-000000000000",
      "ibeacon_major": 178,
      "ibeacon_minor": 40,
      "ibeacon_uuid": "0000149a-0000-0000-0000-000000000000",
      "is_asset": false,
      "mac": "mac4",
      "map_id": "00001148-0000-0000-0000-000000000000",
      "org_id": "00000dbc-0000-0000-0000-000000000000",
      "rssi": 58.22,
      "site_id": "0000245a-0000-0000-0000-000000000000",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "ap_loc": [
        26.98,
        26.97,
        26.96
      ],
      "beam": 124,
      "device_id": "0000254a-0000-0000-0000-000000000000",
      "ibeacon_major": 178,
      "ibeacon_minor": 40,
      "ibeacon_uuid": "0000149a-0000-0000-0000-000000000000",
      "is_asset": false,
      "mac": "mac4",
      "map_id": "00001148-0000-0000-0000-000000000000",
      "org_id": "00000dbc-0000-0000-0000-000000000000",
      "rssi": 58.22,
      "site_id": "0000245a-0000-0000-0000-000000000000",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "ap_loc": [
        26.98,
        26.97,
        26.96
      ],
      "beam": 124,
      "device_id": "0000254a-0000-0000-0000-000000000000",
      "ibeacon_major": 178,
      "ibeacon_minor": 40,
      "ibeacon_uuid": "0000149a-0000-0000-0000-000000000000",
      "is_asset": false,
      "mac": "mac4",
      "map_id": "00001148-0000-0000-0000-000000000000",
      "org_id": "00000dbc-0000-0000-0000-000000000000",
      "rssi": 58.22,
      "site_id": "0000245a-0000-0000-0000-000000000000",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "topic": "topic2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

