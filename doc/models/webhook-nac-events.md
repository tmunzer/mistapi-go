
# Webhook Nac Events

## Structure

`WebhookNacEvents`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookNacEventsEvent`](../../doc/models/webhook-nac-events-event.md) | Optional | - |
| `Topic` | `*string` | Optional | **Default**: `"nac_events"` |

## Example (as JSON)

```json
{
  "topic": "nac_events",
  "events": [
    {
      "ap": "ap6",
      "auth_type": "auth_type0",
      "bssid": "bssid4",
      "dryrun_nacrule_id": "0000114c-0000-0000-0000-000000000000",
      "dryrun_nacrule_matched": false
    },
    {
      "ap": "ap6",
      "auth_type": "auth_type0",
      "bssid": "bssid4",
      "dryrun_nacrule_id": "0000114c-0000-0000-0000-000000000000",
      "dryrun_nacrule_matched": false
    },
    {
      "ap": "ap6",
      "auth_type": "auth_type0",
      "bssid": "bssid4",
      "dryrun_nacrule_id": "0000114c-0000-0000-0000-000000000000",
      "dryrun_nacrule_matched": false
    }
  ]
}
```

