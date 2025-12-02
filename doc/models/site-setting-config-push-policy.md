
# Site Setting Config Push Policy

Mist also uses some heuristic rules to prevent destructive configs from being pushed

## Structure

`SiteSettingConfigPushPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NoPush` | `*bool` | Optional | Stop any new config from being pushed to the device<br><br>**Default**: `false` |
| `PushWindow` | [`*models.PushPolicyPushWindow`](../../doc/models/push-policy-push-window.md) | Optional | If enabled, new config will only be pushed to device within the specified time window |

## Example (as JSON)

```json
{
  "no_push": false,
  "push_window": {
    "enabled": false,
    "hours": {
      "fri": "fri2",
      "mon": "mon8",
      "sat": "sat0",
      "sun": "sun6",
      "thu": "thu6"
    }
  }
}
```

