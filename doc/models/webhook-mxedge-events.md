
# Webhook Mxedge Events

Sample of the `mxedge-events` webhook payload.

## Structure

`WebhookMxedgeEvents`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.MxedgeEvent`](../../doc/models/mxedge-event.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | **Value**: `"mxedge-events"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "audit_id": "53f10664-3ce8-4c27-b382-0ef66432349f",
      "component": "PS1",
      "mxcluster_id": "2815c917-58e7-472f-a190-bfd44fb58d05",
      "mxedge_id": "00000000-0000-0000-1000-020000dc585c",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "service": "tunterm",
      "type": "ME_SERVICE_STOPPED",
      "device_type": "device_type0",
      "from_version": "from_version2",
      "mac": "mac4"
    }
  ],
  "topic": "mxedge-events"
}
```

