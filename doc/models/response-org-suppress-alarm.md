
# Response Org Suppress Alarm

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseOrgSuppressAlarm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Results` | [`[]models.ResponseOrgSuppressAlarmItem`](../../doc/models/response-org-suppress-alarm-item.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "results": [
    {
      "duration": 202,
      "expire_time": 238,
      "scheduled_time": 52,
      "scope": "org",
      "site_id": "00001420-0000-0000-0000-000000000000",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "duration": 202,
      "expire_time": 238,
      "scheduled_time": 52,
      "scope": "org",
      "site_id": "00001420-0000-0000-0000-000000000000",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "duration": 202,
      "expire_time": 238,
      "scheduled_time": 52,
      "scope": "org",
      "site_id": "00001420-0000-0000-0000-000000000000",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

