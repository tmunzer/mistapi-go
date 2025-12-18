
# Webhook Device Events

Sample of the `device-events` webhook payload.

## Structure

`WebhookDeviceEvents`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.DeviceEvent`](../../doc/models/device-event.md) | Required | **Constraints**: *Unique Items Required* |
| `Topic` | `string` | Required, Constant | enum: `device-events`<br><br>**Value**: `"device-events"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "audit_id": "53f10664-3ce8-4c27-b382-0ef66432349f",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "timestamp": 188.18,
      "type": "type0",
      "ap": "ap6",
      "ap_name": "ap_name8",
      "apfw": "apfw6",
      "bandwidth": 64
    }
  ],
  "topic": "device-events"
}
```

