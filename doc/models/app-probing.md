
# App Probing

*This model accepts additional fields of type interface{}.*

## Structure

`AppProbing`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Apps` | `[]string` | Optional | APp-keys from [List Applications](/#operations/listApplications) |
| `CustomApps` | [`[]models.AppProbingCustomApp`](../../doc/models/app-probing-custom-app.md) | Optional | - |
| `Enabled` | `*bool` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
      "name": "name8",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "address": "address4",
      "app_type": "app_type2",
      "hostnames": [
        "hostnames4",
        "hostnames5"
      ],
      "key": "key8",
      "name": "name8",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "enabled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

