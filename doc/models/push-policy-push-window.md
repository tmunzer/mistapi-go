
# Push Policy Push Window

If enabled, new config will only be pushed to device within the specified time window

## Structure

`PushPolicyPushWindow`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Hours` | [`*models.Hours`](../../doc/models/hours.md) | Optional | Days/Hours of operation filter, the available days (mon, tue, wed, thu, fri, sat, sun) |

## Example (as JSON)

```json
{
  "enabled": false,
  "hours": {
    "fri": "fri2",
    "mon": "mon8",
    "sat": "sat0",
    "sun": "sun6",
    "thu": "thu6"
  }
}
```

