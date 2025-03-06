
# Webhook Ping Event

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookPingEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organization |
| `Name` | `string` | Required | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `Timestamp` | `float64` | Required | Epoch (seconds) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "name6",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 38.34,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

