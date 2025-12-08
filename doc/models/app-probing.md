
# App Probing

## Structure

`AppProbing`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Apps` | `[]string` | Optional | APp-keys from [List Applications](../../doc/controllers/constants-definitions.md#list-applications) |
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
      "hostnames": [
        "hostnames4",
        "hostnames5"
      ],
      "key": "key8",
      "name": "name8"
    },
    {
      "address": "address4",
      "app_type": "app_type2",
      "hostnames": [
        "hostnames4",
        "hostnames5"
      ],
      "key": "key8",
      "name": "name8"
    }
  ],
  "enabled": false
}
```

