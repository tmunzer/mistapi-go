
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
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA |
| `OrgId` | `uuid.UUID` | Required | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `Ssids` | `[]string` | Optional | - |
| `Timestamp` | `int` | Required | - |
| `Type` | `string` | Required | event type |
| `Update` | `*bool` | Optional | If presents, represents that this is an update to event with given id sent earlier. default=false |

## Example (as JSON)

```json
{
  "id": "00000afc-0000-0000-0000-000000000000",
  "last_seen": 145.82,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 122,
  "type": "type8",
  "aps": [
    "aps3"
  ],
  "bssids": [
    "bssids8",
    "bssids7"
  ],
  "count": 112,
  "event_id": "000013f4-0000-0000-0000-000000000000",
  "for_site": false
}
```

