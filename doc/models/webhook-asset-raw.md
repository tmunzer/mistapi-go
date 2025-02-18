
# Webhook Asset Raw

Asset raw webhook sample

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookAssetRaw`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookAssetRawEvent`](../../doc/models/webhook-asset-raw-event.md) | Required | List of events<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required | Topic subscribed to<br>**Default**: `"asset-raw"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "events": [
    {
      "asset_id": "00001e56-0000-0000-0000-000000000000",
      "beam": 124,
      "device_id": "0000254a-0000-0000-0000-000000000000",
      "mac": "mac4",
      "map_id": "00001148-0000-0000-0000-000000000000",
      "mfg_company_id": 66.34,
      "mfg_data": "mfg_data2",
      "rssi": 58.22,
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "timestamp": 188.18,
      "ibeacon_major": 178,
      "ibeacon_minor": 40,
      "ibeacon_uuid": "0000149a-0000-0000-0000-000000000000",
      "service_data_data": "service_data_data8",
      "service_data_last_rx_time": 210,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "topic": "asset-raw",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

