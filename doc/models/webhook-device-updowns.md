
# Webhook Device Updowns

device up/down webhook sample

## Structure

`WebhookDeviceUpdowns`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookDeviceUpdownsEvent`](../../doc/models/webhook-device-updowns-event.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required | **Default**: `"device-updowns"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "ap": "ap6",
      "ap_name": "ap_name8",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "site_name": "site_name2",
      "timestamp": 188.18,
      "type": "type0",
      "for_site": false
    }
  ],
  "topic": "device-updowns"
}
```

