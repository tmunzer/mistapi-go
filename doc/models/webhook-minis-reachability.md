
# Webhook Minis Reachability

Sample of the `minis-reachability` webhook payload.

## Structure

`WebhookMinisReachability`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookMinisReachabilityEvent`](../../doc/models/webhook-minis-reachability-event.md) | Required | - |
| `Topic` | `string` | Required, Constant | enum: `minis-reachability`<br><br>**Value**: `"minis-reachability"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "device_mac": "7cb68d8f0440",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "probe_name": "google ping",
      "probe_target": "google.com",
      "probe_type": "reachability",
      "protocol": "icmp",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "test_type": "ping",
      "vlan": 12,
      "avg_latency": 11.76,
      "loss_percentage": 105.92,
      "max_latency": 79.28,
      "min_latency": 246.38
    }
  ],
  "topic": "minis-reachability"
}
```

