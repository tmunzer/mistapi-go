
# App Probing

## Structure

`AppProbing`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Apps` | `[]string` | Optional | app-keys from /api/v1/const/applications |
| `CustomApps` | [`[]models.AppProbingCustomApp`](../../doc/models/app-probing-custom-app.md) | Optional | - |
| `Enabled` | `*bool` | Optional | - |

## Example (as JSON)

```json
{
  "apps": [
    "facebook"
  ],
  "custom_apps": [
    {
      "address": "address4",
      "app_type": "app_type2",
      "hostname": [
        "hostname4",
        "hostname5",
        "hostname6"
      ],
      "name": "name8",
      "network": "network6"
    },
    {
      "address": "address4",
      "app_type": "app_type2",
      "hostname": [
        "hostname4",
        "hostname5",
        "hostname6"
      ],
      "name": "name8",
      "network": "network6"
    }
  ],
  "enabled": false
}
```

