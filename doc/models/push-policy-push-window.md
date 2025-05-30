
# Push Policy Push Window

If enabled, new config will only be pushed to device within the specified time window

*This model accepts additional fields of type interface{}.*

## Structure

`PushPolicyPushWindow`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Hours` | [`*models.Hours`](../../doc/models/hours.md) | Optional | Days/Hours of operation filter, the available days (mon, tue, wed, thu, fri, sat, sun) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "hours": {
    "fri": "fri2",
    "mon": "mon8",
    "sat": "sat0",
    "sun": "sun6",
    "thu": "thu6",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

