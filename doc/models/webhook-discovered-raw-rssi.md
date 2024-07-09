
# Webhook Discovered Raw Rssi

## Structure

`WebhookDiscoveredRawRssi`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookDiscoveredRawRssiEvent`](../../doc/models/webhook-discovered-raw-rssi-event.md) | Optional | - |
| `Topic` | `string` | Required | - |

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
      "site_id": "0000245a-0000-0000-0000-000000000000"
    }
  ],
  "topic": "topic4"
}
```

