
# Webhook Minis Application

Sample of the `minis-application` webhook payload.

## Structure

`WebhookMinisApplication`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookMinisApplicationEvent`](../../doc/models/webhook-minis-application-event.md) | Required | - |
| `Topic` | `string` | Required, Constant | enum: `minis-application`<br><br>**Value**: `"minis-application"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "probe_name": "connectivitycheck.gstatic.com",
      "probe_type": "application",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "test_type": "icmp",
      "vlan": 12,
      "device_mac": "device_mac4",
      "ip": "ip4",
      "latency": 60
    }
  ],
  "topic": "minis-application"
}
```

