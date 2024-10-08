
# Webhook Alarm Event

## Structure

`WebhookAlarmEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aps` | `[]string` | Optional | - |
| `Bssids` | `[]string` | Optional | - |
| `Count` | `*int` | Optional | If present, represents number of events of given type occurred in current interval, default=1 |
| `EventId` | `*uuid.UUID` | Optional | event id |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `uuid.UUID` | Required | - |
| `LastSeen` | `float64` | Required | - |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |
| `OrgId` | `uuid.UUID` | Required | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `Ssids` | `[]string` | Optional | - |
| `Timestamp` | `int` | Required | - |
| `Type` | `string` | Required | event type |
| `Update` | `*bool` | Optional | If presents, represents that this is an update to event with given id sent earlier. default=false |

## Example (as JSON)

```json
{
  "id": "00001c58-0000-0000-0000-000000000000",
  "last_seen": 246.26,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 182,
  "type": "type4",
  "aps": [
    "aps7",
    "aps8",
    "aps9"
  ],
  "bssids": [
    "bssids8",
    "bssids9",
    "bssids0"
  ],
  "count": 52,
  "event_id": "00002550-0000-0000-0000-000000000000",
  "for_site": false
}
```

