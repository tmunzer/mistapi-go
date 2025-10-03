
# Webhook Alarm Event

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookAlarmEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aps` | `[]string` | Optional | - |
| `Bssids` | `[]string` | Optional | - |
| `Count` | `*int` | Optional | If present, represents number of events of given type occurred in current interval, default=1 |
| `EventId` | `*uuid.UUID` | Optional | Event id |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organization |
| `LastSeen` | `models.Optional[float64]` | Optional | Last seen timestamp |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |
| `OrgId` | `uuid.UUID` | Required | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `Ssids` | `[]string` | Optional | - |
| `Timestamp` | `float64` | Required | Epoch (seconds) |
| `Type` | `string` | Required | Event type |
| `Update` | `*bool` | Optional | If presents, represents that this is an update to event with given id sent earlier. default=false |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "last_seen": 1470417522,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 83.74,
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
  "for_site": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

